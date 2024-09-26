package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLink(ctx *gin.Context) {
	var (
		links    []entity.Link
		cacheKey = redis_util.LINK_CACHE_KEY
	)
	if ok := redis_util.GetObject(cacheKey, &links); ok == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(links))
		return
	}
	links = entity.GetLink()
	go redis_util.SetObject(cacheKey, links, redis_util.LINK_CACHE_EXPIRE)
	ctx.JSON(http.StatusOK, respose.Sucess(links))

}
