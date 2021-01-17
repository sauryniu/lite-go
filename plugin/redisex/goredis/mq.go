package goredis

import (
	"github.com/ahl5esoft/lite-go/plugin/mq"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
)

type redisMQ struct {
	Redis redisex.IRedis
}

func (m redisMQ) Publish(channel string, message interface{}) error {
	_, err := m.Redis.Publish(channel, message)
	return err
}

func (m redisMQ) Subscribe(channel string, message chan<- string) {
	m.Redis.Subscribe([]string{channel}, func(sub interface{}) {
		for {
			select {
			case msg := <-sub.(*redis.PubSub).Channel():
				message <- msg.Payload
			default:
			}
		}
	})
}

// NewMQ is 创建reis消息队列
func NewMQ(r redisex.IRedis) mq.IMQ {
	return &redisMQ{
		Redis: r,
	}
}
