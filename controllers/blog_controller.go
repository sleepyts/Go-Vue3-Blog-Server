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
