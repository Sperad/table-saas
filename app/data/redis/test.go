package redis

import (
	redis2 "table-saas/app/components/redis"
)

func Test1() string  {
	val := redis2.Get("a")
	if val == "" {
		val = "123123"
		redis2.Set("a", val, 120)
	}
	return val
}