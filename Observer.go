package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Observer struct {
	C        net.Conn
	B        *bufio.Reader
	VersionS string
	VersionI int64
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
	//Server sets his own Version
	self.readCurrentVersion()
	if err != nil {
		log.Fatal(err)
	}
	ObserverVersion, err := strconv.ParseInt(strings.Trim(version, "\n"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if ObserverVersion < self.VersionI {
		//send observer current version
		buf := bytes.NewBufferString(self.VersionS + "\n")
		_, err := self.C.Write(buf.Bytes())
		if err != nil {
			log.Println(err)
		}
		return true
	}
	//observer is up to date
	buf := bytes.NewBufferString("0\n")
	_, err = self.C.Write(buf.Bytes())
	if err != nil {
		log.Println(err)
	}
	return false
}

func (self *Observer) Monitor() error {
	//Todo:: method to monitor observer	
	log.Println("start monitoring")
	for {
		time.Sleep(time.Second * 5)
	}
	return nil
}

func (self *Observer) ObserverBootstrap() error {
	//if bootstrap needs to be done, observer sends 1 in case he is ready
	//0 if some strange things happened
	stateString, err := self.B.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	state, err := strconv.ParseInt(strings.Trim(stateString, "\n"), 10, 64)
	if state == 1 {
		//send new version
		buf, err := ioutil.ReadFile("client_bin/" + self.VersionS + "/Client")
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

func (self *Observer) readCurrentVersion() {
	self.VersionS = "0"
	self.VersionI = 0
	files, err := filepath.Glob("client_bin/*")
	if err != nil {
		log.Println(err)
	} else if len(files) > 0 {
		self.VersionS = string(filepath.Base(files[len(files)-1]))
		self.VersionI, _ = strconv.ParseInt(filepath.Base(files[len(files)-1]), 10, 64)
	} else {
		log.Println("files empty")
	}
}
