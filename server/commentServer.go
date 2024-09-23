package server

import (
	"Go-Vue3-Blog-Server/globalVar"
	"Go-Vue3-Blog-Server/models/dto"
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetCommentByPage(blogId int, page int) gin.H {
	var (
		pageSize    int
		total       int64
		comments    []entity.Comment
		commentsVOs []vo.CommentVO
		cacheKey    string
	)
	pageSize = 8
	globalVar.Db.Model(entity.Comment{}).Where("blog_id = ?", blogId).Count(&total)

	if page <= 0 || (page-1)*pageSize > int(total) {
		return respose.ErrorWithMsg("Invalid page")
	}
	cacheKey = redis_util.Key2(redis_util.COMMENT_PAGE_CACHE_KEY, blogId, page)
	if cache, err := redis_util.GetObject(cacheKey); err == nil {
		return respose.Sucess(cache)
	}
	globalVar.Db.Limit(pageSize).Offset((page-1)*pageSize).Where("blog_id = ?", blogId).Order("create_time desc").Find(&comments)
	commentsVOs = fromEntityToVO(comments)

	res := respose.NewPageRes(total, commentsVOs)
	go redis_util.SetObject(cacheKey, res, redis_util.COMMENT_PAGE_CACHE_EXPIRE)
	return respose.Sucess(res)
}

func fromEntityToVO(comments []entity.Comment) []vo.CommentVO {
	var commentVOs []vo.CommentVO
	for _, comment := range comments {
		var commentVO vo.CommentVO
		copier.Copy(&commentVO, &comment)
		var replies []entity.Comment
		globalVar.Db.Where("reply_id = ?", comment.Id).Find(&replies)
		commentVO.ReplyList = fromEntityToVO(replies)
		commentVOs = append(commentVOs, commentVO)
	}
	return commentVOs
}

func AddComment(commmentDTO dto.CommentDTO) gin.H {
	// TODO 验证码
	var (
		comment entity.Comment
		key     string
	)
	copier.Copy(&comment, &commmentDTO)
	comment.CreateTime = time.Now()
	if comment.IsAdmin {
		comment.Name = "未月拾叁"
		comment.Url = "https://tsukiyo.cn"
	}
	if comment.Name == "" {
		comment.Name = "匿名"
	}
	result := globalVar.Db.Create(&comment)
	if result.Error != nil {
		fmt.Printf("插入数据失败: %v\n", result.Error)
	}
	go func() {
		key = redis_util.Key1(redis_util.BLOG_COMMENT_PAGE_CACHE_PRE, comment.BlogId)
		redis_util.DeleteKeysWithPrefix(key)
	}()

	return respose.Sucess("")
}

func GetCommentByBlogId(blogId int) gin.H {
	var (
		comments    []entity.Comment
		commentsVOs []vo.CommentVO
	)
	globalVar.Db.Where("blog_id =?", blogId).Order("create_time desc").Find(&comments)
	commentsVOs = fromEntityToVO(comments)
	return respose.Sucess(commentsVOs)
}
