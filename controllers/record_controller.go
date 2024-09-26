package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecordVO(ctx *gin.Context) {
	var (
		cacheKey = redis_util.RECORD_CACHE_KEY
		recordVO []vo.RecordVO
	)
	if err := redis_util.GetObject(cacheKey, &recordVO); err == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(recordVO))
		return
	}
	recordVO = entity.GetRecordVO()
	go redis_util.SetObject(cacheKey, recordVO, redis_util.RECORD_CACHE_EXPIRE)
	ctx.JSON(http.StatusOK, respose.Sucess(recordVO))
}
