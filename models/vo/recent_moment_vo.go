package vo

import "time"

type RecentMomentVO struct {
	Content    string     `json:"content"`
	CreateTime *time.Time `json:"createTime"`
}
