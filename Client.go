package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"net"
	//"time"
)

type Client struct {
	C         net.Conn
	InputChan chan *Input
	B         *bufio.Reader
}

func (self *Client) handleConn() error {
	/*
		for {
			time.Sleep(time.Second * 5)
			log.Println("handle conn dummy func")
		}
	*/
	self.Monitor()
	return nil
}

func (self *Client) Monitor() error {
	for {
		dec := gob.NewDecoder(self.B)
		var i Input
		err := dec.Decode(&i)
		//connection closed
		if err != nil {
			return err
		} else {
			log.Println(i)
			self.InputChan <- &i
		}
	}
	return nil
}

func (self *Client) TestFunc() {
	//Dummy Function for db io stuff
	for i := range self.InputChan {
		if i.Host == "blah" {
			log.Println("blubb")
		} else {
			log.Println(i.Host)
		}
	}
}
