package redis

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"table-saas/config"
)

var redisPool *redis.Pool

var once sync.Once

func GetRedisPool() *redis.Pool {
	once.Do(func() {
		conf := config.Redis
		redisPool = &redis.Pool {
			MaxIdle: conf.MaxIdle,
			MaxActive: conf.MaxActive,
			IdleTimeout: conf.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.Dial(conf.NetWork,conf.Address, redis.DialPassword(conf.Password))
			},
		}
	})
	return redisPool
}