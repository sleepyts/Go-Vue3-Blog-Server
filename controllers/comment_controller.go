package controllers

import (
	"Go-Vue3-Blog-Server/models/dto"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommentByPage(ctx *gin.Context) {
	blogId, err := strconv.Atoi(ctx.Param("blogId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid blog"))
		return
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid page"))
		return
	}
	ctx.JSON(http.StatusOK, server.GetCommentByPage(blogId, page))
}

func GetCommentByBlogId(ctx *gin.Context) {
	blogId, err := strconv.Atoi(ctx.Param("blogId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid blog"))
		return
	}
	ctx.JSON(http.StatusOK, server.GetCommentByBlogId(blogId))
}

func AddComment(ctx *gin.Context) {
	var (
		commentDTO dto.CommentDTO
	)
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid comment"))
		return
	}
	ctx.JSON(http.StatusOK, server.AddComment(commentDTO))
}
