package model

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		log.Println("net listen error, info: ", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener accept error, info: ", err) // err = EOF client退出了
			continue
		}
		// 连接上客户端之后，另开一个线程去处理它的一些业务
		go s.Handler(conn)
	}
}

func (s *Server) Handler(conn net.Conn) {
	// 业务
}
