package entity

import (
	"Go-Vue3-Blog-Server/config"
	"Go-Vue3-Blog-Server/globalVar"
	"Go-Vue3-Blog-Server/models/vo"

	"Go-Vue3-Blog-Server/utils/redis_util"
	"time"

	"github.com/jinzhu/copier"
)

type Comment struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Url        string    `json:"url"`
	BlogId     int       `json:"blogId"`
	ReplyId    int       `json:"replyId"`
	IsAdmin    bool      `json:"isAdmin"`
	CreateTime time.Time `json:"createTime"`
}

func (Comment) TableName() string {
	return "tb_comment"
}
func DeleteCommentById(id uint) error {
	err := globalVar.Db.Where("id = ?", id).Delete(&Comment{}).Error
	// 删除评论缓存
	go func() {
		key := redis_util.Key1(redis_util.BLOG_COMMENT_PAGE_CACHE_PRE, getBlogIdByCommentId(id))
		redis_util.DeleteKeysWithPrefix(key)
	}()
	return err
}
func getBlogIdByCommentId(commentId uint) int {
	var comment Comment
	err := globalVar.Db.Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		return -2
	}
	return comment.BlogId
}
func AddComment(comment *Comment) error {
	comment.CreateTime = time.Now()
	if comment.IsAdmin {
		comment.Name = config.AppConfig.App.Name
		comment.Url = config.AppConfig.App.IndexUrl
	}
	if comment.Name == "" {
		comment.Name = "匿名"
	}
	err := globalVar.Db.Create(comment).Error
	// 删除评论缓存
	go func() {
		key := redis_util.Key1(redis_util.BLOG_COMMENT_PAGE_CACHE_PRE, comment.BlogId)
		redis_util.DeleteKeysWithPrefix(key)
	}()
	return err
}
func GetCommentByBlogId(blogId int) ([]Comment, error) {
	var comments []Comment
	err := globalVar.Db.Where("blog_id = ?", blogId).Order("create_time desc").Find(&comments).Error
	return comments, err
}
func GetCommentByBlogIdAndPage(blogId int, page int, pageSize int) ([]Comment, error) {
	var comments []Comment
	err := globalVar.Db.Order("create_time desc").Limit(pageSize).Offset((page-1)*pageSize).Where("blog_id = ?", blogId).Find(&comments).Error
	return comments, err
}

func GetCountByBlogId(blogId int) (int64, error) {
	var count int64
	err := globalVar.Db.Model(&Comment{}).Where("blog_id = ?", blogId).Count(&count).Error
	return count, err
}

func FromEntityToVO(comments []Comment) []vo.CommentVO {
	var commentVOs []vo.CommentVO
	for _, comment := range comments {
		var commentVO vo.CommentVO
		copier.Copy(&commentVO, &comment)
		var replies []Comment
		globalVar.Db.Where("reply_id = ?", comment.Id).Find(&replies)
		commentVO.ReplyList = FromEntityToVO(replies)
		commentVOs = append(commentVOs, commentVO)
	}
	return commentVOs
}
