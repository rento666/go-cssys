package main

import (
	"github.com/rento666/go-cssys/app/client"
	"github.com/rento666/go-cssys/config"
)

func main() {
	client.Start(config.NetWork, config.IP, config.Port)
}
