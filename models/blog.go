package model

import (
	"time"
)

type Blog struct {
	Id          uint
	Title       string `gorm: unique`
	Img         string
	Content     string
	Decsription string
	CategoryId  uint
	CreateTime  *time.Time
	UpdateTime  *time.Time
}

func (Blog) TableName() string{
	return "tb_blog"
}