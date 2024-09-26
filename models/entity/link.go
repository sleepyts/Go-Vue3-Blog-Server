package entity

import (
	"Go-Vue3-Blog-Server/globalVar"
	"time"
)

type Link struct {
	Id          uint       `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	Description string     `json:"description"`
	Avatar      string     `json:"avatar"`
	CreateTime  *time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime  *time.Time `json:"update_time" gorm:"autoUpdateTime"`
}

func (Link) TableName() string {
	return "tb_friend"
}

func GetLink() []Link {
	var links []Link
	globalVar.Db.Find(&links)
	return links
}
