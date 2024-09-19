package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return func(ctx *gin.Context){
		start:=time.Now()
		ctx.Next()
		log.Print(time.Since(start))
	}
}