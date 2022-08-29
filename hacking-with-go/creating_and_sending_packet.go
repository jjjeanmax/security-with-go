package main

import (
	"log"
	"net"
	"os"
	"time"

	//Get gopacket module
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers" // creating the byte structure for the
	"github.com/google/gopacket/pcap"   // esay way to send a byte
)

var (
	snapshotLen int32 = 1024
	promiscuous       = false
	err         error
	timeout     = -1 * time.Second // show immediately
	handle      *pcap.Handle
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions
)

func main() {
	if len(os.Args) != 2 {
		log.Println("You missed Device Name")
		log.Println("Example: go run port_scaning.go your_device_name_to_scan")
		os.Exit(1)
	}
	deviceName := os.Args[1]

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal("Error openning device. ", err)
	}
	defer handle.Close()
	payload := " Our payload "

	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0x00, 0x00, 0x27, 0x51, 0x1c, 0x5c},
		DstMAC: net.HardwareAddr{0x00, 0x00, 0x27, 0x24, 0xfd, 0x11},
	}

	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{8, 8, 8, 8},
	}

	// TCP layer struct has boolean fields for the SYN, FIN and ACK flags
	// Good for the manipulation and fuzzing TCP handshakes, sessions, and port scan
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(1337),
		DstPort: layers.TCPPort(80),
	}

	// And create the packet with the layers
	buffer := gopacket.NewSerializeBuffer()

	gopacket.SerializeLayers(
		buffer,
		options,
		ethernetLayer,
		ipLayer,
		tcpLayer,
		gopacket.Payload(payload),
	)
	outgoingPacket := buffer.Bytes()

	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal("Error sending packet to network device. ", err)
	}
}
