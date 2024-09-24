package server

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	. "Go-Vue3-Blog-Server/utils/redis_util"

	"github.com/gin-gonic/gin"
)

func GetBlogByPage(page int) gin.H {
	var (
		blogVOs  []vo.BlogVO
		total    int64
		pageSize int
		cacheKey string
	)

	cacheKey = Key1(BLOG_PAGE_CACHE_KEY, page)
	if err := GetObject(cacheKey, &blogVOs); err == nil {
		return respose.Sucess(blogVOs)
	}
	total = entity.GetBlogTotalCount()
	pageSize = 5
	if page <= 0 || (page-1)*pageSize >= int(total) {
		return respose.ErrorWithMsg("Invalid page")
	}
	blogVOs = entity.GetBlogByPage(page, pageSize)
	res := respose.NewPageRes(total, blogVOs)
	go SetObject(cacheKey, res, BLOG_PAGE_CACHE_EXPIRE)
	return respose.Sucess(res)
}

func GetBlogByCategoryIdAndPage(categoryId int, page int) gin.H {
	var (
		blogVOs  []vo.BlogVO
		total    int64
		pageSize int
		cacheKey string
	)

	cacheKey = Key2(BLOG_PAGE_CATEGORY_CACHE_KEY, categoryId, page)
	if err := GetObject(cacheKey, &blogVOs); err == nil {
		return respose.Sucess(blogVOs)
	}
	total = entity.GetBlogCountByCategoryId(categoryId)
	pageSize = 5
	if page <= 0 || (page-1)*pageSize >= int(total) {
		return respose.ErrorWithMsg("Invalid page")
	}
	blogVOs = entity.GetBlogByCategoryIdAndPage(categoryId, page, pageSize)
	res := respose.NewPageRes(total, blogVOs)
	go SetObject(cacheKey, res, BLOG_PAGE_CACHE_EXPIRE)
	return respose.Sucess(res)
}
