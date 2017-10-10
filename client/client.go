package main

import (
	"fmt"
	"net"
)

func main() {

	//Connect udp
	conn, err := net.Dial("udp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("err")
	}
	defer conn.Close()

	//simple write
	conn.Write([]byte("Hello from client"))

	//simple Read
	buffer := make([]byte, 1024)
	conn.Read(buffer)
}
