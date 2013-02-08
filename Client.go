package main

import (
	"bufio"
	"encoding/gob"
	"net"
)

type Client struct {
	Conn      net.Conn
	InputChan chan *Input
	Buff      *bufio.Reader
    Cache     *Cache
}

func (self *Client) handleConn() error {
	return self.Monitor()
}

func (self *Client) Monitor() error {
	for {
        //incoming data from this client in inputchan
		dec := gob.NewDecoder(self.Buff)
		var i Input
		err := dec.Decode(&i)
		if err != nil {
			return err
		} else {
			self.InputChan <- &i
		}
	}
	return nil
}

func (self *Client) ProcessInput() {
	for i := range self.InputChan {
        self.Cache.Insert( i.Time.Second(), i.Time.Minute(), i.Host, i.Value )
	}
}
