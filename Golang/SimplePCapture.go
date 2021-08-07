/* This program does only capturing packets and shows captured ones in raw format */
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Global variables
var (
	snapshotLen int32 = 1024
	promisc     bool  = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func main() {
	//Choosing iface
	var netDevice string
	fmt.Printf("[*] Enter network interface to sniff: ")
	fmt.Scanln(&netDevice)

	// Open device
	fmt.Printf("[*] Packet capturing starts on %s\n", netDevice)
	handle, err := pcap.OpenLive(netDevice, snapshotLen, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use handle as packet source
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for pack := range packetSource.Packets() {
		fmt.Println(pack)
	}
}
