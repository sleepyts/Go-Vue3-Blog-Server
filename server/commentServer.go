package server

import (
	"Go-Vue3-Blog-Server/models/dto"
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"Go-Vue3-Blog-Server/utils/redis_util"

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
	if page <= 0 || (page-1)*pageSize > int(total) {
		return respose.ErrorWithMsg("Invalid page")
	}
	cacheKey = redis_util.Key2(redis_util.COMMENT_PAGE_CACHE_KEY, blogId, page)
	if err := redis_util.GetObject(cacheKey, &commentsVOs); err == nil {
		return respose.Sucess(commentsVOs)
	}
	comments, _ = entity.GetCommentByBlogIdAndPage(blogId, page, pageSize)
	total, _ = entity.GetCountByBlogId(blogId)
	commentsVOs = entity.FromEntityToVO(comments)

	res := respose.NewPageRes(total, commentsVOs)
	go redis_util.SetObject(cacheKey, res, redis_util.COMMENT_PAGE_CACHE_EXPIRE)
	return respose.Sucess(res)
}

func AddComment(commmentDTO dto.CommentDTO) gin.H {
	var (
		comment entity.Comment
	)
	if res := commmentDTO.Verify.Verify(); res != "ok" {
		return respose.ErrorWithMsg(res)
	}
	copier.Copy(&comment, &commmentDTO)

	err := entity.AddComment(&comment)
	if err != nil {
		return respose.ErrorWithMsg(err.Error())
	}

	return respose.Sucess("")
}

func GetCommentByBlogId(blogId int) gin.H {
	var (
		comments    []entity.Comment
		commentsVOs []vo.CommentVO
	)
	comments, _ = entity.GetCommentByBlogId(blogId)
	commentsVOs = entity.FromEntityToVO(comments)
	return respose.Sucess(commentsVOs)
}

func DeleteComment(id int) gin.H {
	err := entity.DeleteCommentById((uint(id)))
	if err != nil {
		return respose.ErrorWithMsg(err.Error())
	}
	return respose.Sucess("")
}
