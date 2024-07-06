package configs

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func InitCache() {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
