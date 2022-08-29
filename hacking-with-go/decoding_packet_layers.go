package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	//Get gopacket module
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap" //wrapper around libcap
)

var (
	// deviceName        = "lo" // eth0 ...
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
		os.Exit(1)
	}
	deviceName := os.Args[1]

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPV4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	// Let's see if the packet is IP (event though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("protocols: ", ip.Protocol)
		fmt.Println()
	}

	// Let's see if packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)

		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}

	// Iterate over all layers , printing out each layer type
	fmt.Println(" All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}

	// Application Layer contains the payload

	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application Layer or Payload Found.")
		fmt.Printf("%s\n", applicationLayer.Payload())

		// Search for a string inside a payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Print("HTTP found")
		}
	}
}
