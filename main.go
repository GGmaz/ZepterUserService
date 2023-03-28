package main

import (
	"ZepterUserService/startup"
	cfg "ZepterUserService/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
