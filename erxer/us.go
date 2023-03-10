package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("from server: Hello, I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	p := make([]byte, 512)
	addr := net.UDPAddr{
		Port: 7788,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error%v\n", err)
		return
	}
	for {
		n, remoteAdd, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v : %s \n", remoteAdd, p[0:n])
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		go sendResponse(ser, remoteAdd)
	}
}
