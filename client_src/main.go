package main

import (
	//"bytes"
	"encoding/gob"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type input struct {
	Host  string
	ID    int64
	Time  time.Time
	Value float64
}

func signalCatcher() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	for signal := range ch {
		if signal == syscall.SIGHUP {
			log.Println("received SIGHUP exiting...")
			os.Exit(0)
		}
	}
}

func startSignalCatcher() {
	//react to sighup
	go signalCatcher()
}

func main() {
	startSignalCatcher()
	conn, err := net.Dial("tcp", "localhost:9035")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		var moep input
		moep.Host = "testhost"
		moep.ID = 123
		moep.Time = time.Now().Local()
		moep.Value = check()
		enc := gob.NewEncoder(conn) // Will write to network.
		err = enc.Encode(moep)
		if err != nil {
			log.Fatal("encode error:", err)
		}
	}

}
