package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, resp string) {
	_, err := conn.WriteToUDP([]byte(resp), addr)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var err error
	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 4000,
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	for {
		received := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(received)
		if err != nil {
			fmt.Println(err)
			continue
		}
		resp := fmt.Sprintf("UDP server: received %s form %v", received[:n], remoteAddr)
		fmt.Println(resp)
		go sendResponse(conn, remoteAddr, resp)
	}
}
