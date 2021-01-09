package redisex

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/thread"
)

type redisLock struct {
	Redis IRedis

	key     string
	seconds time.Duration
}

func (m *redisLock) Lock(format string, args ...interface{}) (bool, error) {
	m.key = fmt.Sprintf(format, args...)
	return m.Redis.Set(m.key, "", "ex", m.seconds, "nx")
}

func (m *redisLock) SetExpire(seconds time.Duration) thread.ILock {
	m.seconds = seconds
	return m
}

func (m *redisLock) Unlock() {
	m.Redis.Del(m.key)
}

// NewLock is thread.ILock
func NewLock(redis IRedis) thread.ILock {
	return &redisLock{
		Redis: redis,
	}
}
