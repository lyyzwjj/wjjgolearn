package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/lyyzwjj/wjjgolearn/02liwenzhou/113blog/controller"
	"github.com/lyyzwjj/wjjgolearn/02liwenzhou/113blog/dao/db"
)

func main() {
	//1 加载数据库
	router := gin.Default()
	dns := `root:Wzzst310@163.com@tcp(wjjzst.com:3306)/blogger?charset=utf8&parseTime=true`
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	// router.StaticFS("/static", http.Dir("/Users/wjj/go/src/github.com/lyyzwjj/wjjgolearn/02liwenzhou/113blog/static"))
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.GET("/article/new/", controller.NewArticle)
	router.POST("/article/submit/", controller.ArticleSubmit)
	router.GET("/article/detail/", controller.ArticleDetail)
	router.POST("/upload/file/", controller.UploadFile)
	router.GET("/leave/new/", controller.LeaveNew)
	router.GET("/about/me/", controller.AboutMe)
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.POST("/leave/submit/", controller.LeaveSubmit)
	_ = router.Run(":8000")

}
