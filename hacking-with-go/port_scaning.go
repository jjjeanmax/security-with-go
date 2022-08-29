package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// var ipToScan = "127.0.0.1"
var minPort = 1
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
		go testTCPConnection(ipToScan, port, doneChanel) // <-- go threads
		activeThreads++
	}

	// wait for all threads to finish
	for activeThreads > 0 {
		<-doneChanel
		activeThreads--

	}
}

func testTCPConnection(ip string, port int, doneChanel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err != nil {
		log.Printf("Host %s has open port: %d\n", ip, port)
	}
	doneChanel <- true
}
