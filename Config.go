package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ObservPort string
	ClientPort string
	//Hosts      []string
	LogFile string
	//Version int64
}

func (self *Config) load(configFile *string) {
	configData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configData, self)
	if err != nil {
		log.Fatal(err)
	}
	//self.setLog()
}

func (self *Config) setLog() {
	logger, err := os.OpenFile(self.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logger)
	log.SetFlags(5)
}
