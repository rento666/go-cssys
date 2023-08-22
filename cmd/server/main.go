package main

import (
	"github.com/rento666/go-cssys/app/server"
	"github.com/rento666/go-cssys/config"
	"github.com/rento666/go-cssys/web"
)

func main() {
	go func() {
		// 开启gin
		web.Start()
	}()
	server.Start(config.IP, config.Port)
}
