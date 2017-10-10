package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	//Connect udp
	conn, err := net.Dial("udp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("err")
	}
	defer conn.Close()

	//simple write
	conn.Write([]byte("Hello from client 1"))

	time.Sleep(2 * time.Second)

	conn.Write([]byte("Hello from client two"))

	//simple Read
	//buffer := make([]byte, 1024)
	//conn.Read(buffer)
}
