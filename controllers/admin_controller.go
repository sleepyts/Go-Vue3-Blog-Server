package controllers

import (
	model "Gin-Learn/models"
	"Gin-Learn/models/respose"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context){
	var admin model.Admin
	if err := ctx.ShouldBindJSON(&admin) ;err!=nil{
		ctx.JSON(http.StatusBadRequest,respose.ErrorWithMsg(err.Error()))
		return 
	}
	println(admin.Password,admin.Username)
}