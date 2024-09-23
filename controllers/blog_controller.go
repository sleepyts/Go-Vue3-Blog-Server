package controllers

import (
	"Go-Vue3-Blog-Server/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页获取博客列表
func GetBlogByPage(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Param("page"))
	ctx.JSON(http.StatusOK, server.GetBlogByPage(page))
}
