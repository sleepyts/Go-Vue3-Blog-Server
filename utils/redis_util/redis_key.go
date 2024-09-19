package redis_util

import (
	"fmt"
	"time"
)

const (
	BLOG_PAGE_CACHE_KEY    = "blog_page_cache:%d"
	BLOG_PAGE_CACHE_EXPIRE = time.Hour * 24 // 1 day
)

func Key(key string, arg interface{}) string {
	return fmt.Sprintf(key, arg)
}
