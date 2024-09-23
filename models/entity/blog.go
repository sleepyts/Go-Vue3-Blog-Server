package entity

import (
	"time"
)

type Blog struct {
	Id          uint  `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Img         string `json:"img"`
	Content     string `json:"content"`
	Description string `json:"description"`
	CategoryId  uint  `json:"categoryId"`
	CreateTime  *time.Time `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

func (Blog) TableName() string{
	return "tb_blog"
}