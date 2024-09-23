package entity

import (
	"time"
)

type Moment struct {
	Id         uint       `json:"id" gorm:"primary_key"`
	Content    string     `json:"content"`
	Visible    bool       `json:"visible"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
}

func (Moment) TableName() string {
	return "tb_moment"
}
