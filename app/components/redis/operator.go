package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

const (
	lLen  = "llen"
	lPush = "lpush"
	rPush = "rpush"
	rPop  = "rpop"
	set   = "set"
	get   = "get"
	del   = "del"
)

func DoCmd(cmd string, args ...interface{}) (interface{}, error) {
	c := GetRedisPool().Get()
	defer c.Close()
	val, err := c.Do(cmd, args...)
	return val, err
}

var lpushWithLimitedScript = redis.NewScript(1, `
	if redis.call("llen", KEYS[1]) < tonumber(ARGV[1]) 
	then
		return redis.call("lpush", KEYS[1], ARGV[2]) 
	else
		return 0
	end
`)


func LPushWithLimited(key string, value string, limited int) (int, error) {
	c := GetRedisPool().Get()
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)
	val, err := lpushWithLimitedScript.Do(c, key, limited, value)
	intVal, err := redis.Int(val, err)
	if err != nil {
		logrus.Error("redis LPushWithLimited key:" + key + ",err:" + err.Error())
		return 0, err
	}
	if intVal == 0 {
		return 0, errors.New("LPushWithLimited failed, reach limited")
	}
	return intVal, nil
}

func LLen(key string) (int, error) {
	var value interface{}
	var err error
	value, err = DoCmd(lLen, key)
	if err != nil {
		logrus.Error("redis lLen key:" + key + ",err:" + err.Error())
		return -1, err
	}
	if value == nil {
		return 0, nil
	}
	intVal, err := redis.Int(value, err)
	if err != nil {
		logrus.Error("redis lLen key:" + key + ",err:" + err.Error())
		return -1, err
	}
	return intVal, nil
}

func LPush(key string, value string) error {
	var err error
	_, err = DoCmd(lPush, key, value)
	if err != nil {
		return err
	}
	return nil
}

func RPush(key string, value string) error {
	var err error
	_, err = DoCmd(rPush, key, value)
	if err != nil {
		return err
	}
	return nil
}

func RPop(key string) string {
	var value interface{}
	var err error
	value, err = DoCmd(rPop, key)
	if err != nil {
		logrus.Error("redis rPop key:" + key + ",err:" + err.Error())
		return ""
	}
	if value == nil {
		return ""
	}
	str, err := redis.String(value, err)
	if err != nil {
		logrus.Error("redis rPop key:" + key + ",err:" + err.Error())
		return ""
	}
	return str
}

func Get(key string) string {
	var value interface{}
	var err error
	value, err = DoCmd(get, key)
	if err != nil {
		logrus.Error("redis get key:" + key + ",err:" + err.Error())
		return ""
	}
	if value == nil {
		return ""
	}
	str, err := redis.String(value, err)
	if err != nil {
		logrus.Error("redis get key:" + key + ",err:" + err.Error())
		return ""
	}
	return str
}

func Set(key string, value string, ttl int) {
	var err error
	_, err = DoCmd(set, key, value, "EX", ttl)
	if err != nil {
		logrus.Error(fmt.Sprintf("redis set key:%s, value:%s, ttl:%d, err:%s", key, value, ttl, err.Error()))
		return
	}
}

func Del(key string) int64 {
	var value interface{}
	var err error
	value, err = DoCmd(del, key)
	if err != nil {
		logrus.Error("redis del key:" + key + ",err:" + err.Error())
		return 0
	}
	if value == nil {
		return 0
	}
	num, err := redis.Int64(value, err)
	if err != nil {
		logrus.Error("redis del key:" + key + ",err:" + err.Error())
		return 0
	}
	return num
}

func SetNX(key string, value string, ttl int) {
	var err error
	_, err = DoCmd(set, key, value, "EX", ttl, "NX")
	if err != nil {
		logrus.Error(fmt.Sprintf("redis setNX key:%s, value:%s, ttl:%d, err:%s", key, value, ttl, err.Error()))
		return
	}
}