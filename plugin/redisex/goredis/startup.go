package goredis

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
)

// IStartupContext is 启动上下文接口
type IStartupContext interface {
	GetRedisOption() *redis.Options
}

// NewStartupHandler is 启动处理器
func NewStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(IStartupContext); ok {
			redisOption := sCtx.GetRedisOption()
			ioc.Set(
				redisex.IoCKey,
				New(redisOption),
			)
		}

		return nil
	})
}
