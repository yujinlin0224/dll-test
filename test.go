package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	lib, err := syscall.LoadDLL("lib.dll")
	if err != nil {
		panic(err)
	}
	defer lib.Release()
	writeProc, err := lib.FindProc("write")
	if err != nil {
		panic(err)
	}
	for {
		var goStr string
		_, err = fmt.Scanln(&goStr)
		if err != nil {
			panic(err)
		}
		cStr, err := syscall.BytePtrFromString(goStr)
		if err != nil {
			panic(err)
		}
		_, _, _ = writeProc.Call(uintptr(unsafe.Pointer(cStr)))
	}
}
