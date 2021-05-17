package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// g

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	// 1. 创建路由
	// 2. 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 加载页面
	r.LoadHTMLGlob("templates/*")
	// 查询所有图书
	r.GET("/book/list", bookListHandler)
	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new_book.html", gin.H{})
		return
	})
	r.POST("/book/new", newBookHandler)
	r.GET("/book/delete", deleteBookHandler)
	// 3. 监听端口,默认在8080
	// r.Run()
	_ = r.Run(":8000")
}

func deleteBookHandler(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "删除id必传"})
		return
	}
	id64, _ := strconv.ParseInt(id, 10, 64)
	err := deleteBook(id64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
	return
}

func newBookHandler(c *gin.Context) {
	// 声明接收的变量
	var book Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := insertBook(book.Title, book.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
	return
}

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "1", "msg": err})
		return
	}

	// 返回数据
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
	return
}
