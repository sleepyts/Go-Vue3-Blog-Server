package vo

import (
	"time"
)

type RandomBlogVO struct {
	Id         uint       `json:"id"`
	Title      string     `json:"title"`
	CreateTime *time.Time `json:"createTime"`
}
