package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApp(ctx *gin.Context) {
	var (
		apps     []entity.App
		cacheKey = redis_util.APP_CACHE_KEY
	)
	if ok := redis_util.GetObject(cacheKey, &apps); ok == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(apps))
		return
	}
	apps = entity.GetApp()
	go redis_util.SetObject(cacheKey, apps, redis_util.APP_CACHE_EXPIRE)
	ctx.JSON(http.StatusOK, respose.Sucess(apps))
}
