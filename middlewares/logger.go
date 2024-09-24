package middlewares

import (
	"log"
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
