package router

import (
	"Go-Vue3-Blog-Server/controllers"
	"Go-Vue3-Blog-Server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Logger())

	r.GET("/settings")

	blogGroup := r.Group("/Blog")
	commentGroup := r.Group("/comment")
	// 博客相关路由
	{
		blogGroup.GET("/page/:page", controllers.GetBlogByPage)
	}
	// 评论相关路由
	{
		commentGroup.GET("/page/:blogId", controllers.GetCommentByPage)
		commentGroup.GET("/:blogId", controllers.GetCommentByBlogId)
		commentGroup.POST("", controllers.AddComment)
		commentGroup.DELETE("")
	}
	return r
}
