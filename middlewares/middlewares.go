package middlewares

import (
	"Go-Vue3-Blog-Server/models/respose"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		go func() {
			log.Printf("\nIp: %s\nRequest time: User-Agent: %s\nProcessed time: %s\nMethod: %s\nPath: %s\n", ctx.ClientIP(), ctx.Request.UserAgent(), time.Since(start), ctx.Request.Method, ctx.Request.URL.Path)

		}()
	}
}

func UserInfoLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Ip", ctx.ClientIP())
		ctx.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token, ok := ctx.Get("token"); ok {
			log.Println(token)
			ctx.Next()
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, respose.ErrorWithMsg("Unauthorized"))

	}
}
