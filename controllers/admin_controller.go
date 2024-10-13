package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	jwtutil "Go-Vue3-Blog-Server/utils/jwt_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var admin entity.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg(err.Error()))
		return
	}
	if entity.IsContain(admin) {
		token, err := jwtutil.GenToken(int64(admin.Id), admin.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, respose.ErrorWithMsg("生成token失败"))
			return
		}
		ctx.JSON(http.StatusOK, respose.Sucess(token))
	} else {
		ctx.JSON(http.StatusUnauthorized, respose.ErrorWithMsg("用户名或密码错误"))
		return
	}
}
