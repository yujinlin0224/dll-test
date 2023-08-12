package main

import "C"
import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

//export write
func write(cStr *C.char) {
	var err error

	goStr := C.GoString(cStr)

	// Print it to the console
	fmt.Println("print:", goStr)

	// Write it to a file
	file, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	if _, err := file.Write([]byte(goStr + "\n")); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Send it to a UDP server
	clientAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	}
	serverAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 4000,
	}
	conn, err := net.DialUDP("udp", &clientAddr, &serverAddr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(goStr))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	buf := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(4 * time.Second))
	n, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println("warn: UDP server did not respond in time")
			return
		} else {
			fmt.Println(err)
			panic(err)
		}
	}
	response := string(buf[:n])
	fmt.Println("Response of UDP server:", response)
}

func main() {}
