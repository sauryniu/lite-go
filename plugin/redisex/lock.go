package redisex

import (
	"time"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/thread"
)

type redisLock struct {
	redis   IRedis
	seconds time.Duration
}

func (m *redisLock) Lock(key string, options ...thread.LockOption) (func(), error) {
	underscore.Chain(options).Each(func(r thread.LockOption, _ int) {
		r(m)
	})

	ok, err := m.redis.Set(key, "", "ex", m.seconds, "nx")
	if ok && err == nil {
		return func() {
			m.redis.Del(key)
		}, nil
	}

	return nil, err
}

// NewLock is thread.ILock
func NewLock(redis IRedis) thread.ILock {
	return &redisLock{
		redis: redis,
	}
}

// NewExpireLockOption is 过期锁选项
func NewExpireLockOption(seconds time.Duration) thread.LockOption {
	return func(lock thread.ILock) {
		lock.(*redisLock).seconds = seconds
	}
}
