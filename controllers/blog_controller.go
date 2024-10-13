package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/server"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页获取博客列表
func GetBlogByPage(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("page参数错误"))
		return
	}
	ctx.JSON(http.StatusOK, server.GetBlogByPage(page))
}

func GetBlogByCategoryIdAndPage(ctx *gin.Context) {
	category, err := strconv.Atoi(ctx.Param("categoryId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("categoryId参数错误"))
		return
	}
	page, _ := strconv.Atoi(ctx.Query("page"))
	ctx.JSON(http.StatusOK, server.GetBlogByCategoryIdAndPage(category, page))
}

func GetBlogById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("blogId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("blogId参数错误"))
		return
	}
	key := redis_util.Key1(redis_util.BLOG_CONTENT_CACHE_KEY, id)
	var blog entity.Blog
	err = redis_util.GetObject(key, &blog)
	if err == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(blog))
	}
	blog = entity.GetBlogById(uint(id))
	ctx.JSON(http.StatusOK, respose.Sucess(blog))

}

func GetBlogList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, respose.Sucess(entity.GetBlogList()))
}

func UpdateBlog(ctx *gin.Context) {
	var blog entity.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("参数错误"))
		return
	}
	entity.UpdateBlog(&blog)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_PAGE_CACHE_KEY)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_CONTENT_CACHE_KEY)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_PAGE_CATEGORY_CACHE_KEY)
	ctx.JSON(http.StatusOK, respose.Sucess(""))
}

func DeleteBlog(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("blogId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("blogId参数错误"))
		return
	}
	entity.DeleteBlog(uint(id))
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_PAGE_CACHE_KEY)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_CONTENT_CACHE_KEY)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_PAGE_CATEGORY_CACHE_KEY)
	ctx.JSON(http.StatusOK, respose.Sucess(""))
}

func AddBlog(ctx *gin.Context) {
	var blog entity.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("参数错误"))
		return
	}
	entity.AddBlog(&blog)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_PAGE_CACHE_KEY)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_CONTENT_CACHE_KEY)
	go redis_util.DeleteKeysWithPrefix(redis_util.BLOG_PAGE_CATEGORY_CACHE_KEY)
	ctx.JSON(http.StatusOK, respose.Sucess(""))
}
