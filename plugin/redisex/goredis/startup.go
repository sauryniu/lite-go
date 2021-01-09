package goredis

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
)

// StartupContext is 启动上下文
type StartupContext struct {
	RedisOption     *redis.Options
	RedisLockOption *redis.Options
}

// NewStartupHandler is 启动处理器
func NewStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(*StartupContext); ok {
			if sCtx.RedisLockOption != nil {
				redis := New(sCtx.RedisLockOption)
				ioc.Set(
					"lock",
					redisex.NewLock(redis),
				)
			}

			if sCtx.RedisOption != nil {
				ioc.Set(
					"redis",
					New(sCtx.RedisOption),
				)
			}
		}

		return nil
	})
}
