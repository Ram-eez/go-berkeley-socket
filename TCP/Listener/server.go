package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	maxDataSize := 10000
	tmp := make([]byte, 1024)
	data := make([]byte, 0, maxDataSize)

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, tmp[:n]...)

		fmt.Printf("Data : %+v\n", string(data))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
