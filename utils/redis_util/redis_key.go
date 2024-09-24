package redis_util

import (
	"fmt"
	"time"
)

const (
	BLOG_PAGE_CACHE_KEY          = "blog_page_cache:%d"
	BLOG_PAGE_CATEGORY_CACHE_KEY = "blog_page_category_cache:%d:%d"
	COMMENT_PAGE_CACHE_KEY       = "comment_page_cache:%d:%d"
	BLOG_COMMENT_PAGE_CACHE_PRE  = "comment_page_cache:%d"
	CATEGORY_CACHE_KEY           = "category_cache:"
	MOMENT_PAGE_CACHE_KEY        = "moment_cache:%d"
	BLOG_PAGE_CACHE_EXPIRE       = time.Hour * 24  // 1 day
	COMMENT_PAGE_CACHE_EXPIRE    = time.Hour * 24  // 1 day
	CATEGORY_CACHE_EXPIRE        = time.Hour * 24  // 1 day
	MOMENT_PAGE_CACHE_EXPIRE     = time.Hour * 24  // 1 day
	COMMENT_VERIFY_CODE_EXPIRE   = time.Minute * 2 // 2 minutes
)

func Key1(key string, arg interface{}) string {
	return fmt.Sprintf(key, arg)
}

func Key2(key string, arg1, arg2 interface{}) string {
	return fmt.Sprintf(key, arg1, arg2)
}
