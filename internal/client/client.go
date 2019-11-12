//client connection - Represents a thermo device
package client

import (
	"net"
	"time"
	"log"
	"math/rand"
	"thermomatic/internal/common"
)

//establish a client connection with server, log in, commence sending Readings 
func StartClient(data []byte) {
	
	conn, err := net.Dial( common.TCP_CONN_TYPE, common.CONN_HOST +":" +common.CONN_PORT) 
	if err != nil {
		log.Printf("StartClient: Dial error %s ", err.Error())
		return
	}

	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		 log.Printf("StartClient: Write Login error %s ", err.Error())
		 return
	}
	log.Printf( "StartClient: Login %v\n", data)
	
	buf := make([]byte, 40)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case _ = <-ticker.C:
			r := CreateRandomReading(rand)
			_ = r.Encode(buf)
			_, err = conn.Write(buf)
			if err != nil {
				 log.Printf("StartClient: Write error %s ", err.Error())
				 return
			}
		}
    }
}



