package redis_util

import (
	"fmt"
	"time"
)

const (
	BLOG_PAGE_CACHE_KEY         = "blog_page_cache:%d"
	COMMENT_PAGE_CACHE_KEY      = "comment_page_cache:%d:%d"
	BLOG_COMMENT_PAGE_CACHE_PRE = "comment_page_cache:%d"
	BLOG_PAGE_CACHE_EXPIRE      = time.Hour * 24 // 1 day
	COMMENT_PAGE_CACHE_EXPIRE   = time.Hour * 24 // 1 day
)

func Key1(key string, arg interface{}) string {
	return fmt.Sprintf(key, arg)
}

func Key2(key string, arg1, arg2 interface{}) string {
	return fmt.Sprintf(key, arg1, arg2)
}
