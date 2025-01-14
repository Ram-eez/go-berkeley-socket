package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:3000")

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	fmt.Println("Listening on port :3000")

	buf := make([]byte, 1024)

	for {
		n, addr, err := ln.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Recieved %s from %s", []byte(buf[:n]), addr)

		_, err = ln.WriteToUDP([]byte(time.Now().String()), addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
