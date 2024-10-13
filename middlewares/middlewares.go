package middlewares

import (
	"Go-Vue3-Blog-Server/models/respose"
	jwtutil "Go-Vue3-Blog-Server/utils/jwt_util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Ip", ctx.ClientIP())
		ctx.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/admin/login" {
			ctx.Next()
			return
		}
		if token, ok := ctx.Request.Header["Authorization"]; ok {
			userName, err := jwtutil.VerifyToken(token[0])
			if err != nil {
				log.Println(err)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, respose.ErrorWithMsg("Unauthorized"))
				return
			}
			ctx.Set("userName", userName)
			ctx.Next()
			return
		}
		log.Println("Admin Auth Fail")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, respose.ErrorWithMsg("Unauthorized"))

	}
}
