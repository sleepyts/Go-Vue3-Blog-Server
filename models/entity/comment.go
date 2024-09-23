package entity

import "time"

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
