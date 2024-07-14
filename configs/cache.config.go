package configs

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func InitCache() {
	host := ViperEnv.GetString("REDIS_HOST")
	port := ViperEnv.GetString("REDIS_PORT")
	Cache = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   0,
	})

	if err := Cache.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
