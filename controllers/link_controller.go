package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"
	"strconv"

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

func AddLink(ctx *gin.Context) {
	var link entity.Link
	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("参数错误"))
		return
	}
	if err := entity.AddLink(link); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("添加失败"))
		return
	}
	go redis_util.DeleteKey(redis_util.LINK_CACHE_KEY)
	ctx.JSON(http.StatusOK, respose.Sucess(nil))
}

func DeleteLink(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("参数错误"))
		return
	}
	if err := entity.DeleteLink(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("删除失败"))
		return
	}
	go redis_util.DeleteKey(redis_util.LINK_CACHE_KEY)
	ctx.JSON(http.StatusOK, respose.Sucess(nil))
}

func UpdateLink(ctx *gin.Context) {
	var link entity.Link
	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("参数错误"))
		return
	}
	if err := entity.UpdateLink(link); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("更新失败"))
		return
	}
	go redis_util.DeleteKey(redis_util.LINK_CACHE_KEY)
	ctx.JSON(http.StatusOK, respose.Sucess(nil))
}
