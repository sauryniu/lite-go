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
	GetRedisLockOption() *redis.Options
}

// NewStartupHandler is 启动处理器
func NewStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(IStartupContext); ok {
			if sCtx.GetRedisLockOption() != nil {
				redis := New(
					sCtx.GetRedisLockOption(),
				)
				ioc.Set(
					"lock",
					redisex.NewLock(redis),
				)
			}

			redisOption := sCtx.GetRedisOption()
			if redisOption != nil {
				ioc.Set(
					"redis",
					New(redisOption),
				)
			}
		}

		return nil
	})
}
