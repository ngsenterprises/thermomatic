package common

import (
	"time"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "1337"
    TCP_CONN_TYPE = "tcp"
	SIMULATOR_NODES = 10
	LOGIN_TIMEOUT = 1 * time.Second
	READIMG_TIMEOUT = 2 * time.Second
)


func HasLength( data []byte, length int ) ( res bool ) {
	defer func() {
		if r := recover(); r != nil {
			res = false
		}
	}()
	_ = data[length -1]
	return true
}

