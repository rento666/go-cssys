package client

import (
	"fmt"
	"log"
	"net"
)

func Start(network, ip string, port int) {
	conn, err := net.Dial(network, fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Println("net dial error, info: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("client success...", conn.RemoteAddr().String())
}
