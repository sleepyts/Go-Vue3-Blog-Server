package server

import (
	"Go-Vue3-Blog-Server/globalVar"
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	. "Go-Vue3-Blog-Server/utils/redis_util"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetBlogByPage(page int) gin.H {
	var (
		blogs    []entity.Blog
		blogVOs  []vo.BlogVO
		total    int64
		pageSize int
		cacheKey string
		wg       *sync.WaitGroup
	)
	pageSize = 5
	cacheKey = Key1(BLOG_PAGE_CACHE_KEY, page)
	if cache, err := GetObject(cacheKey); err == nil {
		return respose.Sucess(cache)
	}
	globalVar.Db.Model(&entity.Blog{}).Count(&total)
	globalVar.Db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&blogs)

	wg = &sync.WaitGroup{}
	wg.Add(len(blogs))
	for _, blog := range blogs {
		go func(blog entity.Blog) {
			var blogVO vo.BlogVO
			copier.Copy(&blogVO, &blog)
			var commentNum int64
			globalVar.Db.Model(&entity.Comment{}).Where("blog_id = ?", blogVO.Id).Count(&commentNum)
			blogVO.CommentNum = int(commentNum)
			globalVar.Db.Model(&entity.Category{}).Select("name").Where("id = ?", blogVO.CategoryId).First(&blogVO.CategoryName)
			blogVOs = append(blogVOs, blogVO)
			wg.Done()
		}(blog)
	}
	wg.Wait()
	res := respose.NewPageRes(total, blogVOs)
	go SetObject(cacheKey, res, BLOG_PAGE_CACHE_EXPIRE)
	return respose.Sucess(res)
}
