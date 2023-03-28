package main

import (
	"zepter/startup"
	cfg "zepter/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
