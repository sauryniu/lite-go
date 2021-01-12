package goredis

import (
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

type startupContext struct{}

func (m startupContext) GetRedisOption() *redis.Options {
	return &redis.Options{
		Addr: "127.0.0.1:6379",
	}
}

func Test_NewStartupHandler(t *testing.T) {
	err := NewStartupHandler().Handle(
		new(startupContext),
	)
	assert.NoError(t, err)

	defer ioc.Remove(redisex.IoCKey)

	assert.True(
		t,
		ioc.Has(redisex.IoCKey),
	)
}
