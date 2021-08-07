/* This program does enumerating network devices */
package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print info
	fmt.Println("Devices Found!")
	for _, dev := range devices {
		fmt.Printf("\nName: %s\n", dev.Name)
		fmt.Printf("Description: %s\n", dev.Description)
		for _, addr := range dev.Addresses {
			fmt.Printf("[+] IP Address: %s\n", addr.IP)
			fmt.Printf("[+] Subnet mask: %s\n", addr.Netmask)
		}
	}
}
