package goredis

import (
	"fmt"
	"reflect"
	"time"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/pubsub"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
)

type goRedis struct {
	client *redis.Client
	subs   map[string]*redis.PubSub
}

func (m goRedis) Close() error {
	return m.client.Close()
}

func (m goRedis) Del(keys ...string) (int, error) {
	count, err := m.client.Del(keys...).Result()
	return int(count), err
}

func (m goRedis) Exists(key string) (bool, error) {
	num, err := m.client.Exists(key).Result()
	return num > 0, err
}

func (m goRedis) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	res, err := m.client.Eval(script, keys, args...).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return res, nil
}

func (m goRedis) Get(key string) (string, error) {
	res, err := m.client.Get(key).Result()
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
	count, err := m.client.Publish(channel, msg).Result()
	return int(count), err
}

func (m goRedis) Set(key, value string, extraArgs ...interface{}) (ok bool, err error) {
	var res string
	if len(extraArgs) == 0 {
		res, err = m.client.Set(key, value, 0).Result()
		ok = res == "OK"
	} else if len(extraArgs) == 1 {
		if extraArgs[0] == "nx" {
			ok, err = m.client.SetNX(key, value, 0).Result()
		} else {
			ok, err = m.client.SetXX(key, value, 0).Result()
		}
	} else if len(extraArgs) == 2 {
		res, err = m.client.Set(
			key,
			value,
			extraArgs[1].(time.Duration),
		).Result()
		ok = res == "OK"
	} else {
		if extraArgs[2] == "nx" {
			ok, err = m.client.SetNX(
				key,
				value,
				extraArgs[1].(time.Duration),
			).Result()
		} else {
			ok, err = m.client.SetXX(
				key,
				value,
				extraArgs[1].(time.Duration),
			).Result()
		}
	}
	return
}

func (m goRedis) Subscribe(channels []string, message chan<- pubsub.Message) {
	sub := m.client.Subscribe(channels...)
	underscore.Chain(channels).Each(func(r string, _ int) {
		m.subs[r] = sub
	})
	go func() {
		for {
			msg := <-sub.Channel()
			message <- pubsub.Message{
				Channel: msg.Channel,
				Text:    msg.Payload,
			}
		}
	}()
}

func (m goRedis) Time() (time.Time, error) {
	return m.client.Time().Result()
}

func (m goRedis) TTL(key string) (time.Duration, error) {
	return m.client.TTL(key).Result()
}

func (m goRedis) Unsubscribe(channels ...string) (err error) {
	underscore.Chain(channels).Map(func(r string, _ int) error {
		if sub, ok := m.subs[r]; ok {
			delete(m.subs, r)
			return sub.Unsubscribe(r)
		}

		return nil
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	return nil
}

// New is IRedis实例
func New(host string, port int, options ...redisex.Option) redisex.IRedis {
	opt := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}
	underscore.Chain(options).Each(func(r redisex.Option, _ int) {
		r(opt)
	})
	return &goRedis{
		client: redis.NewClient(opt),
		subs:   make(map[string]*redis.PubSub),
	}
}
