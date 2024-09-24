package dto

import "Go-Vue3-Blog-Server/models/vo"

type CommentDTO struct {
	Id      uint      `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Content string    `json:"content"`
	Url     string    `json:"url,omitempty"`
	BlogId  int       `json:"blogId"`
	ReplyId int       `json:"replyId"`
	IsAdmin bool      `json:"isAdmin,omitempty"`
	Verify  vo.Verify `json:"verify,omitempty"`
}
