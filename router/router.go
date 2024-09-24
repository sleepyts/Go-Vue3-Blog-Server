package router

import (
	"Go-Vue3-Blog-Server/controllers"
	"Go-Vue3-Blog-Server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Logger())

	r.GET("/settings", controllers.GetSetting)
	r.GET("/verify", controllers.GetCommentVerify)
	blogGroup := r.Group("/Blog")
	commentGroup := r.Group("/comment")
	categoryGroup := r.Group("/category")
	momentGroup := r.Group("/moment")

	// 动态相关路由
	{
		momentGroup.GET("", controllers.GetMomentByPage)
	}
	// 分类相关路由
	{
		categoryGroup.GET("", controllers.GetCategoryList)
	}
	// 博客相关路由
	{
		blogGroup.GET("/page/:page", controllers.GetBlogByPage)
		blogGroup.GET("/category/:categoryId", controllers.GetBlogByCategoryIdAndPage)
	}
	// 评论相关路由
	{
		commentGroup.GET("/page/:blogId", controllers.GetCommentByPage)
		commentGroup.POST("", controllers.AddComment)
		commentGroup.DELETE("")
	}
	return r
}
