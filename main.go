package main

import (
	"flag"
	"runtime"
)

var (
	configFile = flag.String("config", "config/config.json", "main config file")
	mainConfig Config
    cache      *Cache
	server     Server
    update     UpdateServer
    writer     DBWriter
    web        Webserver
)

func init() {
	flag.Parse()
	mainConfig.load(configFile)
    cache = NewCache()
	server.StartServer( cache )
    update.StartUpdateServer()
    writer.init( cache )
    //go web.StartServer( cache )
}

func main() {
	runtime.GOMAXPROCS(4)
    go writer.processInput()
    go update.Run()
	server.Run()
}
