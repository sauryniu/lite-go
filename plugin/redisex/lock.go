package redisex

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/thread"
)

type redisLock struct {
	redis   IRedis
	seconds time.Duration
}

func (m *redisLock) Lock(format string, args ...interface{}) (func(), error) {
	key := fmt.Sprintf(format, args...)
	ok, err := m.redis.Set(key, "", "ex", m.seconds, "nx")
	if ok && err == nil {
		return func() {
			m.redis.Del(key)
		}, nil
	}

	return nil, err
}

func (m *redisLock) SetExpire(seconds time.Duration) thread.ILock {
	m.seconds = seconds
	return m
}

// NewLock is thread.ILock
func NewLock(redis IRedis) thread.ILock {
	return &redisLock{
		redis: redis,
	}
}
