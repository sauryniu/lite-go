package goredis

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/ahl5esoft/lite-go/thread"
	"github.com/go-redis/redis"
)

// ILockStartupContext is 锁启动上下文接口
type ILockStartupContext interface {
	GetRedisLockOption() *redis.Options
}

// NewLockStartupHandler is 锁启动处理器
func NewLockStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(ILockStartupContext); ok {
			redis := New(
				sCtx.GetRedisLockOption(),
			)
			ioc.Set(
				thread.LockIoCKey,
				redisex.NewLock(redis),
			)
		}

		return nil
	})
}
