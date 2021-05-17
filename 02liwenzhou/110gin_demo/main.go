package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

// gin的helloworld

// 定义接收数据的结构体
type Login struct {
	//  binding:"required"修饰的字段, 若接收为空值,则报错,是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// 启动命令修改 Working directory /Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/110gin_demo
func main() {
	fmt.Println(8 << 20)
	// 1. 创建路由
	// 2. 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()

	// {} 为了代码规范

	// 核心是 gin.New() 也可以直接创建不带中间件的路由
	// r := gin.New()

	// 不同的入参接口
	//helloWorld(r)
	//apiParam(r)
	//urlParam(r)
	//postForm(r)
	//uploadFile(r)
	//uploadFiles(r)

	// gin的router前缀树
	// ginGroup(r)	类似RequestMapping
	// ginRouterTree(r)

	// 入参绑定
	// jsonBinding(r)
	// formBinding(r)
	// uriBinding(r)

	// 出参渲染
	// render(r)
	// htmlRender(r)

	// 重定向
	// redirect(r)

	// 异步执行
	// async(r)

	// 中间件
	// middlewareFunc(r)
	// middlewareExercise(r)

	// cookie 和session
	// cookie(r)

	// authMiddleWareDemo(r)

	// 3. 监听端口,默认在8080
	// r.Run()
	r.Run(":8000")
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过,不再调用后续的函数处理
		c.Abort()
		return
	}
}

func authMiddleWareDemo(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", authMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
}

// cookie缺点
// 1. 不安全,明文
// 2. 增加带宽消耗
// 3. 可以被禁用
// 4. cookie有上限
// cookie  cookie 使用操作
func cookie(r *gin.Engine) {
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			// maxAge int  生效时间 单位为s
			// path string cookie所在目录,
			// domain string 域名
			// secure bool  是否只能通过https访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, false)
		}
		fmt.Printf("cookie的值是: %s\n", cookie)
	})
}

func middlewareExercise(r *gin.Engine) {
	// 注册中间件另外实现
	r.Use(MyTime)
	// {} 为了代码规范
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}

}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(time.Second * 2)
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(time.Second * 1)
}

// 定义中间件
func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时", since)

}
func middlewareFunc(r *gin.Engine) {
	// 注册中间件
	r.Use(Middleware())
	// 或者
	// r.Use(MyTime)
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request", req)
			// 页面接收
			c.JSON(http.StatusOK, gin.H{"request": req})
		})

		// 局部中间件  这里省事 用之前的中间件
		r.GET("/middleware1", Middleware(), func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request", req)
			// 页面接收
			c.JSON(http.StatusOK, gin.H{"request": req})
		})
	}
}

// Middleware 定义中间件
// Middleware 中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量Context的key中,可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t1 := time.Since(t)
		fmt.Println("time:", t1)
	}

}

// async 异步执行  在启动新的goroutine时,不应该使用原文的上下文,必须使用它的只读副本
func async(r *gin.Engine) {
	// 1. 异步
	r.GET("/login_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行:" + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/login_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行:" + c.Request.URL.Path)
	})
}

// redirect 重定向
func redirect(r *gin.Engine) {
	// 支持内部和外部重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

}

// htmlRender html渲染
func htmlRender(r *gin.Engine) {
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终将json将title替换
		c.HTML(200, "index.tmpl", gin.H{"title": "我的标题"})
	})

}

// render golang 渲染
func render(r *gin.Engine) {
	// 1. json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	// 2. 结构体响应
	r.GET("someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// 3. XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})
	// 4. YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})
	// 5. protobuf格式,谷歌开发的高效存储的工具
	// 数组?切片?如果自己构建一个传输格式,应该是什么格式?
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
}

// curl http://127.0.0.1:8000/loginUri/root/admin -X POST
// formBinding Form 绑定
func uriBinding(r *gin.Engine) {
	r.POST("/loginUri/:user/:password", func(c *gin.Context) {
		// 声明接收的变量
		var form Login
		// ShouldBindUri 解析uri中的参数
		// 根据请求头Content-Type 自动推断
		// 将request的uri中的数据,自动解析到结构体
		if err := c.ShouldBindUri(&form); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// curl http://127.0.0.1:8000/loginForm -d 'username=root&password=admin' -X POST
// formBinding Form 绑定
func formBinding(r *gin.Engine) {
	r.POST("/loginForm", func(c *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头Content-Type 自动推断
		// 将request的表单中的数据,自动解析到结构体
		if err := c.Bind(&form); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// curl http://127.0.0.1:8000/loginJSON -H 'Content-Type:application/json' -d '{"user":"root","password":"admin"}' -X POST
// jsonBinding JSON 绑定
func jsonBinding(r *gin.Engine) {
	r.POST("/loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据,自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
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
