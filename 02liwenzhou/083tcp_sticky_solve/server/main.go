package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/lyyzwjj/wjjgolearn/02liwenzhou/083tcp_sticky_solve/protocol"
)

// socket_stick/server/main.go
// Nagle算法  客户端攒了很多tcp包 攒到一定程度再发给server端   客户端没收到上一次发送的包的ack 并且本次发的包很小 就会有Nagle 等到大于MSS
// 只允许有一个未被ack的包
/*
	若是包长度达到MSS，则容许发送
	若是包含FIN，则容许发送
	若是设置了TCP_NODELAY，则容许发送
	未设置TCP_CORK选项时，若全部发出去的小数据包（包长度小于MSS）均被确认，则容许发送
	上述条件都未知足，但发生了超时（通常为200ms），则当即发送。
*/

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// var buf [1024]byte
	for {
		//n, err := reader.Read(buf[:])
		//if err == io.EOF {
		//	break
		//}
		//if err != nil {
		//	fmt.Println("read from client failed, err:", err)
		//	break
		//}
		//recvStr := string(buf[:n])
		recvStr, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode failed, err", err)
			return
		}
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
