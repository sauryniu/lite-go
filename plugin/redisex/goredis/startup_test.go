package goredis

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/thread"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func Test_NewStartupHandler_RedisLockOption(t *testing.T) {
	err := NewStartupHandler().Handle(&StartupContext{
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
	err := NewStartupHandler().Handle(&StartupContext{
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
