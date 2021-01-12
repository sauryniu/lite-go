package goredis

import (
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/thread"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

type lockStartupContext struct{}

func (m lockStartupContext) GetRedisLockOption() *redis.Options {
	return &redis.Options{
		Addr: "127.0.0.1:6379",
	}
}

func Test_NewLockStartupHandler(t *testing.T) {
	err := NewLockStartupHandler().Handle(
		new(lockStartupContext),
	)
	assert.NoError(t, err)

	defer ioc.Remove(thread.LockIoCKey)

	assert.True(
		t,
		ioc.Has(thread.LockIoCKey),
	)
}
