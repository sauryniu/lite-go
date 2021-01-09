package goredis

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/thread"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

type startupContext struct {
	RedisLockOption *redis.Options
	RedisOption     *redis.Options
}

func (m startupContext) GetRedisLockOption() *redis.Options {
	return m.RedisLockOption
}

func (m startupContext) GetRedisOption() *redis.Options {
	return m.RedisOption
}

func Test_NewStartupHandler_RedisLockOption(t *testing.T) {
	err := NewStartupHandler().Handle(&startupContext{
		RedisLockOption: &redis.Options{
			Addr: "127.0.0.1:6379",
		},
	})
	assert.NoError(t, err)

	defer ioc.Remove("lock")

	lock := ioc.Get("lock").(thread.ILock)

	duration := 1 * time.Second
	ok, err := lock.SetExpire(duration).Lock("test")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = lock.SetExpire(duration).Lock("test")
	assert.NoError(t, err)
	assert.False(t, ok)
}

func Test_NewStartupHandler_RedisOption(t *testing.T) {
	err := NewStartupHandler().Handle(&startupContext{
		RedisOption: &redis.Options{
			Addr: "127.0.0.1:6379",
		},
	})
	assert.NoError(t, err)

	defer ioc.Remove("redis")

	assert.True(
		t,
		ioc.Has("redis"),
	)
}
