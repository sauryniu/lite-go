package goredis

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
)

type goRedis struct {
	Client *redis.Client
}

func (m goRedis) Get(key string) (string, error) {
	res, err := m.Client.Get(key).Result()
	if err != redis.Nil {
		return "", err
	}

	return res, nil
}

func (m goRedis) Del(keys ...string) (int, error) {
	count, err := m.Client.Del(keys...).Result()
	return int(count), err
}

func (m goRedis) Exists(key string) (bool, error) {
	num, err := m.Client.Exists(key).Result()
	return num > 0, err
}

func (m goRedis) Set(key, value string, extraArgs ...interface{}) (ok bool, err error) {
	var res string
	if len(extraArgs) == 0 {
		res, err = m.Client.Set(key, value, 0).Result()
		ok = res == "OK"
	} else if len(extraArgs) == 1 {
		if extraArgs[0] == "nx" {
			ok, err = m.Client.SetNX(key, value, 0).Result()
		} else {
			ok, err = m.Client.SetXX(key, value, 0).Result()
		}
	} else if len(extraArgs) == 2 {
		fmt.Println(extraArgs[1])
		res, err = m.Client.Set(
			key,
			value,
			extraArgs[1].(time.Duration),
		).Result()
		ok = res == "OK"
	} else {
		if extraArgs[2] == "nx" {
			ok, err = m.Client.SetNX(
				key,
				value,
				extraArgs[1].(time.Duration),
			).Result()
		} else {
			ok, err = m.Client.SetXX(
				key,
				value,
				extraArgs[1].(time.Duration),
			).Result()
		}
	}
	return
}

func (m goRedis) SetNX(key, value string, expires time.Duration) (bool, error) {
	return m.Client.SetNX(key, value, expires).Result()
}

func (m goRedis) Time() (time.Time, error) {
	return m.Client.Time().Result()
}

// New is IRedis实例
func New(options *redis.Options) redisex.IRedis {
	return &goRedis{
		Client: redis.NewClient(options),
	}
}
