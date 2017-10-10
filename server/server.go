package main

import (
	"fmt"
	"net"
)

type Address struct {
	transport string
	ip        string
}

func (addr Address) Network() string { // name of the network (for example, "tcp", "udp")
	return addr.transport
	//String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")

}

func (addr Address) String() string { // name of the network (for example, "tcp", "udp")
	return addr.ip
	//String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")

}

func handleConnection(c *net.UDPConn) {
	//Simple read from connection
	buffer := make([]byte, 1024)
	//c.Read(buffer)
	//fmt.Println("output:", string(buffer))
	n, addr, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("msg:", buffer)
	}
	fmt.Println("n:", n)
	fmt.Println("addr:", addr)
	fmt.Println("output:", string(buffer))

	//msg := string(buf[0:n])
	//m := server.parseMessage(msg)
}

func main() {
	address := "127.0.0.1"
	port := "4000"
	udpAddress, err := net.ResolveUDPAddr("udp4", address+":"+port)

	pc, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("err:", err)
	}
	defer pc.Close()

	for {
		handleConnection(pc)
	}

	//simple read
	//buffer := make([]byte, 1024)
	//pc.ReadFrom(buffer)
	//fmt.Println("output:", string(buffer))

	//simple write
	//pc.WriteTo([]byte("Hello from client"), addr)
}
