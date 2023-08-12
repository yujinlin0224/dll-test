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
	serverAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 4000,
	}
	conn, err := net.ListenUDP("udp", &serverAddr)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	for {
		n, fromAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		received := string(buf[:n])
		response := fmt.Sprintf("received %s form %v", received, fromAddr)
		fmt.Println("UDP server:", response)
		go sendResponse(conn, fromAddr, response)
	}
}
