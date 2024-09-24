package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryList(ctx *gin.Context) {
	var categoryList []entity.Category
	err := redis_util.GetObject(redis_util.CATEGORY_CACHE_KEY, &categoryList)
	if err == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(categoryList))
		return
	}
	categorys, err := entity.GetCategoryList()
	res := respose.Sucess(categorys)
	if err != nil {
		ctx.JSON(http.StatusOK, respose.ErrorWithMsg(err.Error()))
		return
	}
	go redis_util.SetObject(redis_util.CATEGORY_CACHE_KEY, categorys, redis_util.CATEGORY_CACHE_EXPIRE)
	ctx.JSON(http.StatusOK, res)
}
