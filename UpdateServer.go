package main

import (
    "bufio"
    "log"
    "net"
)

type UpdateServer struct {
    updateListener net.Listener
    ToStart chan *Observer
}

func (self *UpdateServer) StartUpdateServer() {
    var err error
    self.updateListener, err = net.Listen("tcp", ":" + mainConfig.ObservPort)
    if err != nil {
        log.Fatal("Create updateServer listen port: " + err.Error())
    }
    self.ToStart = make(chan *Observer, 1000)
    go self.Accept()
}

func (self *UpdateServer) Accept() {
    for {
		conn, err := self.updateListener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		observer := &Observer{
            C: conn,
            B: bufio.NewReader(conn),
        }
		self.ToStart <- observer
		log.Println("added observer for " + observer.C.RemoteAddr().String())
	}
}

func (self *UpdateServer) startUpdate( observer *Observer ) {
    log.Println("started observer for " + observer.C.RemoteAddr().String())
	err := observer.HandleConn()
	//handle closed connection
	//make sure failed connections return errors!
	if err != nil {
		log.Println("ObserverHandler returned:" + err.Error())
		observer.C.Close()
	}
}

func (self *UpdateServer) Run() {
    for observer := range self.ToStart {
        go self.startUpdate( observer )
    }
}
