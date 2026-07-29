package service

import (
	"time"

	"github.com/kikudesuyo/buildlog/api/cache"
)

const contentCacheTTL = 30 * time.Second

var contentCache = cache.New(contentCacheTTL)
