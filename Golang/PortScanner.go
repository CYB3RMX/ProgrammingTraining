/* This program does scan all 65535 TCP ports */
/* You can change the Port number anything you want */
package main

import (
    "fmt"
    "net"
)

func main(){
    fmt.Printf("[*] Enter target host: ")

    var target string

    fmt.Scanln(&target)

    fmt.Printf("\n[*] Scanning: %s please wait...\n", target)
    fmt.Println("+ ------------------------------ +\n")
    for i := 1; i <= 65535; i++ {
        address := fmt.Sprintf("%s:%d", target, i)
        conn, err := net.Dial("tcp", address)
        if err != nil {
            continue
        }
        conn.Close()
        fmt.Printf("[+] Port: %d is open.\n", i)
    }
    fmt.Println("\n+ ------------------------------ +")
}
