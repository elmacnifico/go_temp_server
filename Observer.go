package main

import (
	"bufio"
	"bytes"
	//"encoding/gob"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type Observer struct {
	C net.Conn
	B *bufio.Reader
}

func (self *Observer) HandleConn() error {
	if self.VersionOutdated() {
		err := self.ObserverBootstrap()
		if err != nil {
			log.Println(err)
		}
	}
	return self.Monitor()
}

func (self *Observer) VersionOutdated() bool {
	//Observer sends first his version
	version, err := self.B.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ObserverVersion, err := strconv.ParseInt(strings.Trim(version, "\n"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if ObserverVersion < mainConfig.Version {
		//Tell Observer he is not up to date
		buf := bytes.NewBufferString("0\n")
		_, err := self.C.Write(buf.Bytes())
		if err != nil {
			log.Println(err)
		}
		return true
	}
	//observer is up to date
	buf := bytes.NewBufferString("1\n")
	_, err = self.C.Write(buf.Bytes())
	if err != nil {
		log.Println(err)
	}
	return false
}

func (self *Observer) Monitor() error {
	//Todo:: method to monitor observer	
	/*

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
	*/
	log.Println("start monitoring")
	for {
		//TODO
		//Read version from config yaml
		//check if version exists
		//send version to client

		buf := bytes.NewBufferString("test\n")
		_, err := self.C.Write(buf.Bytes())
		if err != nil {
			self.C.Close()
			return err
		}
		time.Sleep(time.Second * 5)
	}
	return nil
}

func (self *Observer) ObserverBootstrap() error {
	//if bootstrap needs to be done, observer send 1 in case he is ready
	//0 is some strange things happened
	stateString, err := self.B.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	state, err := strconv.ParseInt(strings.Trim(stateString, "\n"), 10, 64)
	if state == 1 {
		//send new version
		buf, err := ioutil.ReadFile("client_bin/Client")
		if err != nil {
			log.Println(err)
		}
		n, err := self.C.Write(bytes.NewBufferString(strconv.Itoa(binary.Size(buf)) + "\n").Bytes())
		if err != nil {
			log.Println(err)
		}
		n, err = self.C.Write(buf)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(n)
		}
		log.Println("send new version")
	} else {
		//handle this crazy error
		return errors.New("Observer not ready for update")
	}
	return nil
}

/*
func (self *Observer) TestFunc() {
	//Dummy Function for db io stuff
	for i := range self.InputChan {
		if i.Host == "blah" {
			log.Println("blubb")
		}
	}
}
*/
