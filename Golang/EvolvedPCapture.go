/* Same as SimplePCapture.go but this one does packet decoding */
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Setting variables
var (
	snapshotLen int32 = 1024
	promisc     bool  = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func main() {
	// Choosing interface to sniff
	var netDevice string
	fmt.Printf("[*] Enter network interface to sniff: ")
	fmt.Scanln(&netDevice)

	// Accessing network device
	handle, err := pcap.OpenLive(netDevice, snapshotLen, promisc, timeout)

	// error handling
	if err != nil {
		log.Fatalln(err)
	}
	defer handle.Close()

	// handling packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Printf("[*] Packet capturing starts on %s\n", netDevice)
	for packet := range packetSource.Packets() {
		packetInfo(packet)
	}
}

// parsing packet informations
func packetInfo(packet gopacket.Packet) {
	// parsing IP layer
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		// getting layer variables
		ips, _ := ipLayer.(*layers.IPv4)

		// printing
		fmt.Printf("\nSrcAddr: %s  => DstAddr: %s\n", ips.SrcIP, ips.DstIP)
	}

	// parsing TCP layer
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		// getting layer variables
		tcp, _ := tcpLayer.(*layers.TCP)

		// printing
		fmt.Printf("SrcPort: %d  => DstPort: %d\n", tcp.SrcPort, tcp.DstPort)
	}
}
