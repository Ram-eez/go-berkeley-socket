package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// tmp := make([]byte, 1024)
	data := make([]byte, 1024)

	for {
		// n, err := conn.Read(tmp)
		_, err := conn.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		// data = append(data, tmp[:n]...)

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
