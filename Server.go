package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

type Server struct {
	L_client      net.Listener
	ToStartClient chan *Client
    Cache         *Cache
}

func (self *Server) StartServer( cache *Cache ) {
	var err error
	self.L_client, err = net.Listen("tcp", ":"+mainConfig.ClientPort)
	if err != nil {
		log.Fatal("Create client listen port: " + err.Error())
	}
	self.ToStartClient = make(chan *Client, 1000)
    self.Cache = cache

	go self.AcceptClient()
}

func (self *Server) AcceptClient() {
	for {
		c, err := self.L_client.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := &Client{
            Conn: c,
            InputChan: make(chan *Input, 100000),
            Buff: bufio.NewReader(c),
            Cache: self.Cache,
        }
		self.ToStartClient <- client
		log.Println("added client for " + client.Conn.RemoteAddr().String())
	}
}

func (self *Server) startClient(client *Client) {
	log.Println("started client for " + client.Conn.RemoteAddr().String())
	err := client.handleConn()
	if err != nil {
		log.Println("ClientHandler returned: " + err.Error())
		client.Conn.Close()
	}
}

func (self *Server) ClientStarter() {
	for client := range self.ToStartClient {
		go client.ProcessInput()
		go self.startClient(client)
	}
}

func (self *Server) Run() {
	go self.ClientStarter()
	for {
		//replace this with function which checks for new source
		// builds new binaries and set new versions
		time.Sleep(time.Second * 10)
	}
}
