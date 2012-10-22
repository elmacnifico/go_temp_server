package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

type Server struct {
	L_observ      net.Listener
	L_client      net.Listener
	ToStartObserv chan *Observer
	ToStartClient chan *Client
}

func (self *Server) StartServer() {
	var err error
	self.L_observ, err = net.Listen("tcp", ":"+mainConfig.ObservPort)
	if err != nil {
		log.Fatal("Create observer listen port: " + err.Error())
	}
	self.L_client, err = net.Listen("tcp", ":"+mainConfig.ClientPort)
	if err != nil {
		log.Fatal("Create observer listen port: " + err.Error())
	}
	self.ToStartObserv = make(chan *Observer, 1000)
	self.ToStartClient = make(chan *Client, 1000)

	go self.AcceptObserver()
	go self.AcceptClient()
}

func (self *Server) AcceptObserver() {
	for {
		c, err := self.L_observ.Accept()
		if err != nil {
			log.Fatal(err)
		}
		observer := &Observer{C: c, B: bufio.NewReader(c)}
		self.ToStartObserv <- observer
		log.Println("added observer for " + observer.C.RemoteAddr().String())
	}
}

func (self *Server) AcceptClient() {
	for {
		c, err := self.L_client.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := &Client{C: c, InputChan: make(chan *Input, 100000), B: bufio.NewReader(c)}
		self.ToStartClient <- client
		log.Println("added client for " + client.C.RemoteAddr().String())
	}
}

func (self *Server) startObserver(observer *Observer) {
	log.Println("started observer for " + observer.C.RemoteAddr().String())
	err := observer.HandleConn()
	//handle closed connection
	//make sure failed connections return errors!
	if err != nil {
		log.Println("ObserverHandler returned:" + err.Error())
		observer.C.Close()
	}
}

func (self *Server) startClient(client *Client) {
	log.Println("started client for " + client.C.RemoteAddr().String())
	err := client.handleConn()
	if err != nil {
		log.Println("ClientHandler returned: " + err.Error())
		client.C.Close()
	}
}

func (self *Server) ClientStarter() {
	for client := range self.ToStartClient {
		go client.TestFunc()
		go self.startClient(client)
	}
}

func (self *Server) ObserverStarter() {
	for observer := range self.ToStartObserv {
		go self.startObserver(observer)
	}
}

func (self *Server) Run() {
	go self.ClientStarter()
	go self.ObserverStarter()
	for {
		//replace this with function which checks for new source
		// builds new binaries and set new versions
		time.Sleep(time.Second * 10)
	}

}
