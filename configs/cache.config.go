package configs

import (
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func InitCache() {
	Cache = redis.NewClient(&redis.Options{
		Addr: ViperEnv.Get("REDIS_HOST").(string) + ":" + ViperEnv.Get("REDIS_PORT").(string),
		DB:   0,
	})
}
