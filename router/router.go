package router

import (
	"Go-Vue3-Blog-Server/controllers"
	"Go-Vue3-Blog-Server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(middlewares.UserInfoLogger())
	r.GET("/settings", controllers.GetSetting)
	r.GET("/verify", controllers.GetCommentVerify)
	r.GET("/record", controllers.GetRecordVO)
	adminGroup := r.Group("/admin")
	blogGroup := r.Group("/Blog")
	commentGroup := r.Group("/comment")
	categoryGroup := r.Group("/category")
	momentGroup := r.Group("/moment")
	appGroup := r.Group("/app")
	linkGroup := r.Group("/links")
	adminGroup.Use(middlewares.AdminAuth())

	// APP相关路由
	{
		appGroup.GET("", controllers.GetApp)
	}
	// 友链相关路由
	{
		linkGroup.GET("", controllers.GetLink)
	}
	// 动态相关路由
	{
		momentGroup.GET("", controllers.GetMomentByPage)
		momentGroup.POST("/like/:momentId", controllers.LikeOrUnlikeMoment)
	}
	// 分类相关路由
	{
		categoryGroup.GET("", controllers.GetCategoryList)
	}
	// 博客相关路由
	{
		blogGroup.GET("/page/:page", controllers.GetBlogByPage)
		blogGroup.GET("/category/:categoryId", controllers.GetBlogByCategoryIdAndPage)
		blogGroup.GET("/:blogId", controllers.GetBlogById)
	}
	// 评论相关路由
	{
		commentGroup.GET("/page/:blogId", controllers.GetCommentByPage)
		commentGroup.POST("", controllers.AddComment)
		commentGroup.DELETE("")
	}
	return r
}
