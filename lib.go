package main

// #include <stdlib.h>
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
	lAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	}
	rAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 4000,
	}
	conn, err := net.DialUDP("udp", &lAddr, &rAddr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	conn.SetReadBuffer(4096)
	conn.SetWriteBuffer(4096)
	_, err = conn.Write([]byte(goStr))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	resp := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(4 * time.Second))
	_, err = bufio.NewReader(conn).Read(resp)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println("warn: UDP server did not respond in time")
		} else {
			fmt.Println(err)
			panic(err)
		}
	} else {
		fmt.Println("response:", string(resp))
	}
}

func main() {}
