package config

import (
	"time"
)

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Extra    string
	PoolSize int
	Attempts int
	Interval time.Duration
}

var MysqlConf = &MysqlConfig {
	Host: "127.0.0.1",
	User: "root",
	Password: "root",
	Port: 9306,
	Database: "table_engine",
	Extra: "charset=utf8&parseTime=true&interpolateParams=true&autocommit=true&loc=Local",
	PoolSize: 4,
	Attempts: 3,
	Interval: 10,
}