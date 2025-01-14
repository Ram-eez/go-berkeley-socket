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

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		_, err := conn.Write([]byte("Hello I am yes\n"))
		if err != nil {
			log.Fatal(err)
		}

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Read %s from %s", []byte(buf[:n]), addr)

		time.Sleep(time.Second)
	}
}
