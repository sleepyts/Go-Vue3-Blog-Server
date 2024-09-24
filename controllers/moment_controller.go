package controllers

import (
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMomentByPage(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Page参数错误"))
		return
	}
	ctx.JSON(http.StatusOK, server.GetMomentByPage(page))
}
