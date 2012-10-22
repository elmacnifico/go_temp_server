package main

import (
	//"bufio"
	"bytes"
	"encoding/json"
	"flag"
	//"io"
	"io/ioutil"
	"log"
	"net"
	//"os"
	//"os/exec"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Server string
}

var (
	configFile = flag.String("config", "config/config.json", "main config file")
	mainConfig Config
	myVersion  = "0\n"
)

func init() {
	flag.Parse()
	configData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configData, &mainConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func getNewVersion() {
	log.Println("get new version")
}

func main() {
	conn, err := net.Dial("tcp", mainConfig.Server)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 100)
	//send 0 as default version of base client
	conn.Write(bytes.NewBufferString(myVersion).Bytes())
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	} else {
		//resp from buf should always show that version is outdated
		resp, err := strconv.ParseInt(strings.Trim(string(buf[:n]), "\n"), 10, 64)
		if err != nil {
			log.Println(err)
		} else if resp == 0 {
			conn.Write(bytes.NewBufferString("1\n").Bytes())
			getNewVersion()
			buf := make([]byte, 10000000)
			log.Println("going to read from conn")
			readbytes := 0
			for {
				n, err := conn.Read(buf)
				if n == 0 {
					break
				}
				if err != nil {
					log.Println(err.Error() + "read bin")
				}
				readbytes += n
			}
			log.Println("finished read from conn")
			log.Println(readbytes)
			newBinary := buf[:readbytes]
			err = ioutil.WriteFile("new_client/base", newBinary, 0770)
			if err != nil {
				log.Println(err)
			}

		} else {
			log.Println("something went terribly wrong...")
		}
	}
	for {
		log.Println("going to sleep")
		time.Sleep(1000000000)
	}
}
