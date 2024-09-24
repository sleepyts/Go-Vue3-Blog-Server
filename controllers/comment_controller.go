package controllers

import (
	"Go-Vue3-Blog-Server/models/dto"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"Go-Vue3-Blog-Server/server"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"crypto/rand"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCommentByPage(ctx *gin.Context) {
	blogId, err := strconv.Atoi(ctx.Param("blogId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid blog"))
		return
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid page"))
		return
	}
	ctx.JSON(http.StatusOK, server.GetCommentByPage(blogId, page))
}

func GetCommentByBlogId(ctx *gin.Context) {
	blogId, err := strconv.Atoi(ctx.Param("blogId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid blog"))
		return
	}
	ctx.JSON(http.StatusOK, server.GetCommentByBlogId(blogId))
}

func AddComment(ctx *gin.Context) {
	var (
		commentDTO dto.CommentDTO
	)
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, respose.ErrorWithMsg("Invalid comment"))
		return
	}
	ctx.JSON(http.StatusOK, server.AddComment(commentDTO))
}

func GetCommentVerify(ctx *gin.Context) {
	var (
		verifyVO vo.Verify
		key      string
	)
	var1, _ := rand.Int(rand.Reader, big.NewInt(10))
	verifyVO.Var1 = var1.String()
	var2, _ := rand.Int(rand.Reader, big.NewInt(10))
	verifyVO.Var2 = var2.String()
	key = strconv.FormatInt((time.Now().UnixNano() / 1e6), 10)
	verifyVO.Key = key
	verifyVO.VerifyVar = ""
	redis_util.SetObject(key, verifyVO, time.Minute*2)
	ctx.JSON(http.StatusOK, respose.Sucess(verifyVO))
}
