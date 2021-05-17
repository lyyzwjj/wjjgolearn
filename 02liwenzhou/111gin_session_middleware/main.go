package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// g

func main() {
	fmt.Println(8 << 20)
	// 1. 创建路由
	// 2. 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()

	// 3. 监听端口,默认在8080
	// r.Run()
	r.Run(":8000")
}
