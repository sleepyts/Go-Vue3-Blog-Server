package vo

import "time"

type CommentVO struct {
	Id         uint        `json:"id"`
	Name       string      `json:"name"`
	Content    string      `json:"content"`
	Url        string      `json:"url"`
	BlogId     int         `json:"blogId"`
	ReplyId    int         `json:"replyId"`
	IsAdmin    bool        `json:"isAdmin"`
	CreateTime *time.Time  `json:"createTime"`
	ReplyList  []CommentVO `json:"replyList"`
}
