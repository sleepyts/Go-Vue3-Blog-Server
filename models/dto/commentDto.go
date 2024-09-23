package dto

type CommentDTO struct {
	Id      uint   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content"`
	Url     string `json:"url,omitempty"`
	BlogId  int    `json:"blogId"`
	ReplyId int    `json:"replyId"`
	IsAdmin bool   `json:"isAdmin,omitempty"`
}
