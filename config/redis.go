package config

import "time"

type RedisConfig struct {
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
	NetWork string 				`default:"tcp"`
	Address string
	Password string
}

var Redis = &RedisConfig{
	MaxIdle: 10,
	MaxActive: 3,
	IdleTimeout: time.Duration(300) * time.Second,
	NetWork: "tcp",
	Address: "127.0.0.1:6379",
	Password: "sperad-pwd",
}