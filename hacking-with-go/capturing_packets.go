/*
	Equivalent to tcpdump -i eth0 -w network_traffic.pcap
	deviceName        = "lo" // eth0 ...

*/

package main

import (
	"fmt"
	"os"
	"time"

	//Get gopacket module
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap" //wrapper around libcap
	"github.com/google/gopacket/pcapgo"
)

var (
	snapshotLen int32 = 1024
	promiscuous       = false
	err         error
	timeout     = -1 * time.Second // show immediately
	handle      *pcap.Handle
	packetCount = 0
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You missed Device Name")
		fmt.Println("Example: go run port_scaning.go your_device_name_to_scan")
		os.Exit(1)
	}
	deviceName := os.Args[1]

	// Open output pcap file and write header
	f, _ := os.Create("network_traffic.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error openning device %s: %v", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	// start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture 100 packets and then stop
		if packetCount > 100 {
			break
		}
	}
}
