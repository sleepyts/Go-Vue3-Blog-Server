package controllers

import (
	"Gin-Learn/entity/entity"
	"Gin-Learn/entity/respose"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context){
	var admin entity.Admin
	if err := ctx.ShouldBindJSON(&admin) ;err!=nil{
		ctx.JSON(http.StatusBadRequest,respose.ErrorWithMsg(err.Error()))
		return 
	}
	println(admin.Password,admin.Username)
}