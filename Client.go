package main

import (
	"bufio"
	//"log"
	"net"
)

type Client struct {
	C         net.Conn
	InputChan chan *Input
	B         *bufio.Reader
}
