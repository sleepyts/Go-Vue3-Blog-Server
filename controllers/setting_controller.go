package controllers

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetSetting(ctx *gin.Context) {
	var (
		settings vo.SettingsVO
	)
	copier.Copy(&settings, entity.GetSettings())
	settings.RandomBlogVOs = entity.GetRandomBlogVO()
	ctx.JSON(http.StatusOK, respose.Sucess(settings))
}
