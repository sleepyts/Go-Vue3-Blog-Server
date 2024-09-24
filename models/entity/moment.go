package entity

import (
	"Go-Vue3-Blog-Server/globalVar"
	"Go-Vue3-Blog-Server/utils/redis_util"
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

func GetMomentByPage(page, pageSize int) ([]*Moment, error) {
	var (
		moments  []*Moment
		cacheKey = redis_util.Key1(redis_util.MOMENT_PAGE_CACHE_KEY, page)
	)
	err := redis_util.GetObject(cacheKey, &moments)
	if err == nil {
		return moments, nil
	}

	err = globalVar.Db.Order("create_time desc").Offset((page-1)*pageSize).Limit(pageSize).Where("visible =?", 1).Find(&moments).Error
	go redis_util.SetObject(cacheKey, moments, redis_util.MOMENT_PAGE_CACHE_EXPIRE)
	return moments, err
}

func GetMomentCount() (int64, error) {
	var count int64
	err := globalVar.Db.Model(&Moment{}).Where("visible =?", 1).Count(&count).Error
	return count, err
}
