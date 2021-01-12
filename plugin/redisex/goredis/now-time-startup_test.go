package goredis

import (
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/timeex"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

type nowTimeStartupContext struct{}

func (m nowTimeStartupContext) GetRedisNowTimeOption() *redis.Options {
	return &redis.Options{
		Addr: "127.0.0.1:6379",
	}
}

func Test_NewNowTimeStartupHandler(t *testing.T) {
	err := NewNowTimeStartupHandler().Handle(
		new(nowTimeStartupContext),
	)
	assert.NoError(t, err)

	defer ioc.Remove(timeex.NowTimeIoCKey)

	assert.True(
		t,
		ioc.Has(timeex.NowTimeIoCKey),
	)
}
