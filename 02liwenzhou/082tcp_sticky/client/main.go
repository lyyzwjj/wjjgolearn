package main

import (
	"fmt"
	"net"
)

// tcp/client/main.go

// 客户端
// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		msgByte := []byte(msg)
		//head := [4]byte{26, 0, 0, 0}
		//bytes := append(head[:], msgByte...)
		//conn.Write(bytes)
		conn.Write(msgByte)
	}
}
