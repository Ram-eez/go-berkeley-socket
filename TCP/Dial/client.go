package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	var Message string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}
		Message = scanner.Text()
		_, err = conn.Write([]byte(Message))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Message sent")
	}

}
