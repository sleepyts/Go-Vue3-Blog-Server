package vo

import (
	"time"
)

type RecordVO struct {
	Year       string      `json:"year"`
	RecordList []RecordDTO `json:"recordList"`
}

type RecordDTO struct {
	Id         int        `json:"id"`
	Title      string     `json:"title"`
	CreateTime *time.Time `json:"createTime"`
}
