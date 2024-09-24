package vo

import "time"

type BlogVO struct {
	Id           uint       `json:"id" gorm:"primary_key"`
	Title        string     `json:"title"`
	Img          string     `json:"img"`
	Description  string     `json:"description"`
	CategoryId   uint       `json:"categoryId"`
	CreateTime   *time.Time `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime"`
	CommentNum   int        `json:"commentNum"`
	CategoryName string     `json:"categoryName"`
}
