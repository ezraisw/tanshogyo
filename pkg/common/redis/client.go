package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ProvideUniversalClient(configGetter RedisConfigGetter) redis.UniversalClient {
	config := configGetter.GetRedisConfig()

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
}
