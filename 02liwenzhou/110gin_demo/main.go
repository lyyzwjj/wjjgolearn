package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// gin的helloworld
func main() {
	fmt.Println(8 << 20)
	// 1. 创建路由
	// 2. 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 核心是 gin.New() 也可以直接创建不带中间件的路由
	// r := gin.New()

	// 不同的接口
	//helloWorld(r)
	//apiParam(r)
	//urlParam(r)
	//postForm(r)
	//uploadFile(r)
	//uploadFiles(r)
	//ginGroup(r)
	ginRouterTree(r)

	// 3. 监听端口,默认在8080
	// r.Run()
	r.Run(":8000")
}

func ginRouterTree(r *gin.Engine) {
	r.POST("/", login)
	r.POST("search", login)
	r.POST("support", login)
	r.POST("/blog/:post", login)
	r.POST("/contact", login)
	r.POST("/blog/:post", login)
	r.POST("/about", login)
}

// 路由组
// gin 路由使用了httprouter   https://github.com/julienschmidt/httprouter
//
func ginGroup(r *gin.Engine) {
	// 路由组1, 处理GET请求
	// curl http://localhost;8000/v1/login?name=zhangsan
	v1 := r.Group("/v1")
	// {} 书写规范  不写也可以
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	// curl -X POST http://localhost;8000/v2/login?name=zhangsan -d {}
	// curl -H “Content-Type:application/json” -X POST --data ‘[{“index”:["*"],“preference”:“1503652289983”,“ignore_unavailable”:“true”},{“sort”:[{“timestamp”:{“order”:“desc”}}],“query”:{“must_not”:[],“bool”:{“must”:[{“query_string”:{“query”:“cluster”}},{“range”:{“timestamp”:{“gte”:“1503667558137”,“lte”:“1503667558137”}}}]}},“from”:“0”,“size”:“500”,“version”:“true”}]’ http://127.0.0.1:18080/
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "hello %s\n", name)
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, "hello %s\n", name)
}
func uploadFiles(r *gin.Engine) {
	// 限制表单上传大小 8MB 默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/uploadfiles", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有文件
		files := form.File["files"]
		// 遍历所有图片
		dir := "/Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/110gin_demo/"
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, dir+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("uplaod ok %d files", len(files)))

	})
}
func uploadFile(r *gin.Engine) {
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dir := "/Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/110gin_demo/"
		// 传到项目根目录, 名字就用本身的
		_ = c.SaveUploadedFile(file, dir+file.Filename)
		// 打印信息
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})
}

func postForm(r *gin.Engine) {
	r.POST("/form", func(c *gin.Context) {
		// 表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alert")
		// 接受其他的
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 多选框
		hobbies := c.PostFormArray("hobby")
		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, username is %s, password is %s, hobbies is %v",
				type1, username, password, hobbies))
	})
}

func urlParam(r *gin.Engine) {
	// URL参数可以通过DefaultQuery()或Query()方法获取
	// DefaultQuery 参数不存在则 返回默认值 Query()若不存在则返回空串
	// http://localhost:8000/welcom?name=wjj
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "jack") // jack是默认值
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})
}

func apiParam(r *gin.Engine) {
	// 只有: 是硬匹配 name后面不能接 /了
	// r.GET("/user/:name", func(c *gin.Context) {
	// * 是软匹配 取到的东西都会带/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		// http://localhost:8000/user/zhangsan/haha
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action) // zhangsan is /haha
	})
}

func helloWorld(r *gin.Engine) {
	// 2. 绑定路由规则, 执行的函数
	// gin.Context,封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.POST("/xxxPost", getting)
	r.PUT("/xxxPut")
}

//
func getting(c *gin.Context) {

}
