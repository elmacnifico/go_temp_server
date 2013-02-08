package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	oldWBytes float64
	newWBytes float64
	oldRBytes float64
	newRBytes float64
	oldWOps   float64
	newWOps   float64
	oldROps   float64
	newROps   float64
	oldTime   time.Time
	err       error
)

func removeDoubleSpaces(line *string) {
	myRegex, err := regexp.Compile("\\s+")
	if err != nil {
		log.Println("regex: " + err.Error())
	}
	*line = strings.TrimSpace(myRegex.ReplaceAllLiteralString(*line, " "))
}

func rIOOps(Ops string, diff time.Duration) float64 {
	newROps, err := strconv.ParseFloat(Ops, 64)
	if err != nil {
		log.Println("parse float" + err.Error())
	}
	myValue := (oldROps - newROps) / diff.Seconds()
	oldROps = newROps
	return myValue
}

func wIOOps(Ops string, diff time.Duration) float64 {
	newWOps, err := strconv.ParseFloat(Ops, 64)
	if err != nil {
		log.Println("parse float" + err.Error())
	}
	myValue := (oldWOps - newWOps) / diff.Seconds()
	oldWOps = newWOps
	return myValue
}

func rdiskStats(Bytes string, diff time.Duration) float64 {
	newRBytes, err = strconv.ParseFloat(Bytes, 64)
	if err != nil {
		log.Println("parse float: " + err.Error())
	}
	myValue := 512 * (oldRBytes - newRBytes) / diff.Seconds() / 1024 / 1024
	//insert send func here :P
	oldRBytes = newRBytes
	return myValue
}

func wdiskStats(Bytes string, diff time.Duration) float64 {
	newWBytes, err = strconv.ParseFloat(Bytes, 64)
	if err != nil {
		log.Println("parse float: " + err.Error())
	}
	myValue := 512 * (oldWBytes - newWBytes) / diff.Seconds() / 1024 / 1024
	//insert send func here :P
	oldWBytes = newWBytes
	return myValue
}

func analyze(line string, now time.Time) float64 {
	diff := oldTime.Sub(now)
	removeDoubleSpaces(&line)
	values := strings.Split(line, " ")
	//log.Printf("wstat %v\n", wdiskStats(values[9], diff))
	//log.Printf("rstat %v\n", rdiskStats(values[5], diff))
	//log.Printf("rIOOPS: %v\n", rIOOps(values[3], diff))
	//log.Printf("wIOOPS: %v\n", wIOOps(values[7], diff))
	oldTime = now
	return wdiskStats(values[9], diff)
}

func check() float64 {
	myDev := "sdb3"

	file, err := os.Open("/proc/diskstats")
	defer file.Close()
	if err != nil {
		log.Println("open: " + err.Error())
	}
	fileBuf := bufio.NewReader(file)
	time.Sleep(time.Millisecond * 100)
	for {
		myLine, err := fileBuf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("read: " + err.Error())
			}
			break
		} else {
			match, _ := regexp.MatchString(myDev, myLine)
			if match {
				return analyze(myLine, time.Now())
			}
		}
	}

	return 0
}
