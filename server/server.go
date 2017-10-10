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
func main() {
	// listen to incoming udp packets
	//var addr = Address{
	//transport: "udp",
	//ip:        "127.0.0.1",
	//}

	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:4000")
	pc, err := net.ListenPacket("udp", "127.0.0.1:4000")
	if err != nil {
		//log.Fatal(err)
		fmt.Println("err")
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)
	fmt.Println("output:", string(buffer))

	//simple write
	pc.WriteTo([]byte("Hello from client"), addr)
}
