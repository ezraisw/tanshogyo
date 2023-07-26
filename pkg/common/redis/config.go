package redis

type RedisConfigGetter interface {
	GetRedisConfig() *RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}
