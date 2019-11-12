package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"thermomatic/internal/common"
	"thermomatic/internal/server"
	"thermomatic/internal/simulator"
)

func main() {

	go server.StartServer( common.CONN_HOST, common.CONN_PORT )

	time.Sleep(250 * time.Millisecond)

	go simulator.StartSimulator()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	reader.ReadString('\n')

}



