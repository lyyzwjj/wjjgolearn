package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func httpHandler1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("/Users/wjj/go/src/github.com/lyyzwjj/wjjgolearn/02liwenzhou/085http_server/server/test.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func httpHandler2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求,参数都放在URL上(query param) 查询参数, 请求体中是没有数据的
	fmt.Println(r.URL)
	// fmt.Println(r.URL.Query())  // 自动帮我们识别URL中的query param
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 客户端请求的body
	w.Write([]byte("ok"))

}

func main() {
	http.HandleFunc("/hello/", httpHandler1)
	http.HandleFunc("/xxx/", httpHandler2)
	// http.ListenAndServe("127.0.0.1:9090", nil)
	// 所有都能访问到
	http.ListenAndServe("0.0.0.0:9090", nil)

	// 自定义的Server
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        myHandler,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//log.Fatal(s.ListenAndServe())
}
