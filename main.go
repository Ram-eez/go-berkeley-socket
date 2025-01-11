package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	tmp := make([]byte, 1024)
	data := make([]byte, 0)

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, tmp[:n]...)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)

	}
}
