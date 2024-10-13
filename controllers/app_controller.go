package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const cacheKey = redis_util.APP_CACHE_KEY

func GetApp(ctx *gin.Context) {
	var (
		apps []entity.App
	)
	if ok := redis_util.GetObject(cacheKey, &apps); ok == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(apps))
		return
	}
	apps = entity.GetApp()
	go redis_util.SetObject(cacheKey, apps, redis_util.APP_CACHE_EXPIRE)
	ctx.JSON(http.StatusOK, respose.Sucess(apps))
}

func AddApp(ctx *gin.Context) {
	var app entity.App
	if err := ctx.ShouldBindJSON(&app); err == nil {
		entity.AddApp(app)
		ctx.JSON(http.StatusOK, respose.Sucess(nil))
		go redis_util.DeleteKey(cacheKey)
	}

}

func DeleteApp(ctx *gin.Context) {
	id, error := strconv.Atoi(ctx.Param("id"))
	if error != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("id参数错误"))
	}
	entity.DeleteApp((uint(id)))
	go redis_util.DeleteKey(cacheKey)
	ctx.JSON(http.StatusOK, respose.Sucess(nil))
}

func UpdateApp(ctx *gin.Context) {
	var app entity.App
	if err := ctx.ShouldBindJSON(&app); err == nil {
		entity.UpdateApp(app)
		ctx.JSON(http.StatusOK, respose.Sucess(nil))
		go redis_util.DeleteKey(cacheKey)
	}
}
