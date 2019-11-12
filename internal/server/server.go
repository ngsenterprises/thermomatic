package server

import (
	"fmt"
	"os"
	"net"
	"log"
	"time"
	"sync"
	"thermomatic/internal/common"
	"thermomatic/internal/imei"
	"thermomatic/internal/client"
)

type SafeImeiCodes struct {
	V   map[uint64]bool
	Mux sync.Mutex
}

//add imei code to code Map
func (c *SafeImeiCodes) AddCode(key uint64) {
	c.Mux.Lock()
	c.V[key] = true
	c.Mux.Unlock()
}

//delete imei code from code Map
func (c *SafeImeiCodes) DelCode(key uint64) {
	c.Mux.Lock()
	delete(c.V, key)
	c.Mux.Unlock()
}

//check code map if key exists
func (c *SafeImeiCodes) Contains(key uint64) bool {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	_, contains := c.V[key]
	if !contains {
		return false	
	} 
	return true
}

//run as go routine for each device connection
func handleConnection(conn net.Conn, codes SafeImeiCodes ) {
	
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println( "handleConnection: Client connected from " + remoteAddr)        

	defer conn.Close()
	
	// retrieve login imei code
	loginBuf := make([]byte, 15)
	conn.SetReadDeadline(time.Now().Add( common.LOGIN_TIMEOUT ))
	_, err := conn.Read(loginBuf)
	if err != nil {
		log.Printf("handleConnection: login error %s", err.Error())
		return
	}
	conn.SetReadDeadline( time.Time{} )
	log.Printf("handleConnection: login %v", loginBuf)
	
	// validate and convert imei code
	imeiCode, err := imei.Decode(loginBuf)
	if err != nil {
		log.Printf( "handleConnection: iemi error %s", err.Error() )
		return
	} 
	log.Printf("handleConnection: imeiCode %d", imeiCode)

	defer func() {
		codes.DelCode(imeiCode)
	}()

	//check if imei code is already logged in
	loggedIn := codes.Contains(imeiCode)	
	if loggedIn {
		log.Printf("handleConnection: duplicate login %d ", imeiCode)
		return
	}
	codes.AddCode(imeiCode)

	//commence getting Reading's from device
	reading := client.Reading{}
	readingBuf := make([]byte, 40)
	conn.SetReadDeadline(time.Now().Add( common.READIMG_TIMEOUT ))
	for {
		_, err = conn.Read(readingBuf)
		if err != nil {
			log.Printf("handleConnection: device Reading error %s", err.Error())
			return
		}
		reading.Decode(readingBuf)
		fmt.Printf("Reading: remoteAddr %s %v : \n", remoteAddr, reading )
		conn.SetReadDeadline(time.Now().Add( common.READIMG_TIMEOUT ))
	}
}

//start the tcp/ip server  
func StartServer( host string, port string ) error {
	
	listener, err := net.Listen( common.TCP_CONN_TYPE, host+ ":" +port )
	if err != nil {
		log.Printf( "StartServer: error listening: %s ", err.Error() )
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening on " +common.CONN_HOST + ":" +common.CONN_PORT)

	codes := SafeImeiCodes{ V: make(map[uint64]bool) }

	//accept connections and start device processing
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		} 
		
		go handleConnection(conn, codes)
	}
	
}


