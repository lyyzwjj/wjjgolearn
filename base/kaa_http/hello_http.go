package main

import (
	"fmt"
	"net/http"
	"time"
)

/**
 * @author  wjj
 * @date  2020/9/11 1:53 上午
 * @description 简单的web服务器
 * 如果传入/time/xxx  其他的路由地址 走默认的路由规则
 */

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	// 同样请求localhost:8080/time/1
	// http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		// w.Header()
		w.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
	// 不传handler 会使用DefaultServeMux
}
