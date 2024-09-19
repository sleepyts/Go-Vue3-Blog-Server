package router

import (
	"Gin-Learn/middlewares"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine{
	r:=gin.Default()
	r.Use(middlewares.Logger())
	return r;
}