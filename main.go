package main

import (
	"flag"
)

var (
	configFile = flag.String("config", "config/config.json", "main config file")
	mainConfig Config
	server     Server
)

func init() {
	flag.Parse()
	mainConfig.load(configFile)
	server.StartServer()
}

func main() {
	server.Run()
}
