package goredis

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
)

type goRedis struct {
	Client *redis.Client
}

func (m goRedis) Close() error {
	return m.Client.Close()
}

func (m goRedis) Del(keys ...string) (int, error) {
	count, err := m.Client.Del(keys...).Result()
	return int(count), err
}

func (m goRedis) Exists(key string) (bool, error) {
	num, err := m.Client.Exists(key).Result()
	return num > 0, err
}

func (m goRedis) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	res, err := m.Client.Eval(script, keys, args...).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return res, nil
}

func (m goRedis) Get(key string) (string, error) {
	res, err := m.Client.Get(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return res, nil
}

func (m goRedis) Publish(channel string, message interface{}) (int, error) {
	var msg string
	if reflect.TypeOf(message).Kind() == reflect.String {
		msg = message.(string)
	} else {
		msg, _ = jsoniter.MarshalToString(message)
	}
	count, err := m.Client.Publish(channel, msg).Result()
	return int(count), err
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

func (m goRedis) Subscribe(channels []string, handleAction func(sub interface{})) {
	sub := m.Client.Subscribe(channels...)
	go handleAction(sub)
}

func (m goRedis) Time() (time.Time, error) {
	return m.Client.Time().Result()
}

// New is IRedis实例
func New(host string, port int, password string) redisex.IRedis {
	return &goRedis{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: password,
		}),
	}
}
