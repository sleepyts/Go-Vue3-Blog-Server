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
	ip, _ := ctx.Get("Ip")
	ctx.JSON(http.StatusOK, server.GetMomentByPage(page, ip.(string)))
}

func LikeOrUnlikeMoment(ctx *gin.Context) {
	momentId, err := strconv.Atoi(ctx.Param("momentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("MomentId参数错误"))
		return
	}
	ip, ok := ctx.Get("Ip")
	if !ok {
		return
	}
	ctx.JSON(http.StatusOK, server.LikeOrUnlikeMoment(momentId, ip.(string)))
}
