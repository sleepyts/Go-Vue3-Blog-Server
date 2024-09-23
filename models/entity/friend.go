package entity

import "time"

type Friend struct {
	Id          uint       `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	Avatar      string     `json:"avatar"`
	Description string     `json:"description"`
	CreateTime  *time.Time `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

func (Friend) TableName() string {
	return "tb_friend"
}
