package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var minPort = 0
var maxPort = 1024

func main() {
	if len(os.Args) != 2 {
		log.Println("Ip address bad or missed")
		log.Println("Example: go run port_scaning.go your_ip_to_scan")
		os.Exit(1)
	}

	ipToScan := os.Args[1]

	activeThreads := 0
	doneChanel := make(chan bool)

	for port := minPort; port <= maxPort; port++ {
		go grabBanner(ipToScan, port, doneChanel) // <-- go threads
		activeThreads++
	}

	// wait for all threads to finish
	for activeThreads > 0 {
		<-doneChanel
		activeThreads--
	}
}

func grabBanner(ip string, port int, doneChanel chan bool) {
	connection, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err != nil {
		doneChanel <- true
		return
	}
	// see if server offers anything to read
	buffer := make([]byte, 4096)
	connection.SetReadDeadline(time.Now().Add(time.Second * 5))
	// set timeout
	numBytesRead, err := connection.Read(buffer)
	log.Println(numBytesRead)

	if err != nil {
		doneChanel <- true
		return
	} else {
		panic(err)
	}
	log.Printf("Banner from port %d\n%s\n", port, buffer[0:numBytesRead])
	doneChanel <- true
}
