package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 例题：golang实现RPC程序，实现求矩形面积和周长

type Params struct {
	Width, Height int
}

type Rect struct{}

// 结构体字段首字母要大写，可以别人调用
// 函数名必须首字母大写
// 函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型
// 函数还必须有一个返回值error

// Area RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// Perimeter 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

// 主函数
func main() {
	// 1.注册服务
	rect := new(Rect)
	// 注册一个rect的服务
	err := rpc.Register(rect)
	if err != nil {
		return
	}
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
