package controllers

import (
	"Gin-Learn/entity/entity"
	"Gin-Learn/entity/respose"
	"Gin-Learn/entity/vo"
	"Gin-Learn/globalVar"
	. "Gin-Learn/utils/redis_util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// 分页获取博客列表
func GetBlogByPage(ctx *gin.Context) {
	var (
		blogs    []entity.Blog
		blogVOs  []vo.BlogVO
		total    int64
		page     int
		pageSize int
		cacheKey string
	)
	page, _ = strconv.Atoi(ctx.Param("page"))
	pageSize = 5
	cacheKey = Key(BLOG_PAGE_CACHE_KEY, page)
	if cache, err := GetObject(cacheKey); err == nil {
		ctx.JSON(http.StatusOK, respose.Sucess(cache))
		return
	}
	globalVar.Db.Model(&entity.Blog{}).Count(&total)
	globalVar.Db.Limit(5).Offset((page - 1) * pageSize).Find(&blogs)
	blogVOs = make([]vo.BlogVO, len(blogs))
	for index, blog := range blogs {
		var blogVO vo.BlogVO
		if err := copier.Copy(&blogVO, &blog); err != nil {
			return
		}
		blogVO.CommentNum = 123
		blogVO.CategoryName = "分类名称"
		blogVOs[index] = blogVO
	}
	res := respose.NewPageRes(total, blogVOs)
	go SetObject(cacheKey, res, BLOG_PAGE_CACHE_EXPIRE)
	ctx.JSON(http.StatusOK, respose.Sucess(res))
}
