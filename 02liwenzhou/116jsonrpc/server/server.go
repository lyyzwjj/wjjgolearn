package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Arith 结构体，用于注册的
type Arith struct{}

// ArithRequest 声明参数结构体
type ArithRequest struct {
	A, B int
}

// ArithResponse 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

// Multiply 乘法
func (a *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// Divide 商和余数
func (a *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 除
	res.Quo = req.A / req.B
	// 取模
	res.Rem = req.A % req.B
	return nil
}

// 主函数
func main() {
	// 1.注册服务
	rect := new(Arith)
	// 注册一个rect的服务
	err := rpc.Register(rect)
	if err != nil {
		return
	}
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	// 循环监听服务
	for {
		var conn net.Conn
		conn, err = listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new Client")
			jsonrpc.ServeConn(conn)
		}(conn)

	}
}
