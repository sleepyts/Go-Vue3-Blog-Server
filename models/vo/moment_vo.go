package vo

import (
	"time"
)

type MomentVO struct {
	Id         uint       `json:"id" gorm:"primary_key"`
	Content    string     `json:"content"`
	Visible    bool       `json:"visible"`
	LikeCount  uint       `json:"likeCount"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
}
