package server

import "github.com/rento666/go-cssys/app/server/model"

func Start(ip string, port int) {
	server := model.NewServer(ip, port)
	server.Start()
}
