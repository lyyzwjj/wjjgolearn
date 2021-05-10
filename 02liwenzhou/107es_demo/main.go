package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// es insert data demo
// Student ...
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	// 1. 初始化一个client客户端用来连接
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := Student{"Rion", 22, false}
	// 链式操作
	put1, err := client.
		Index().
		Index("student").
		// Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}
