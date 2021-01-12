package goredis

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/ahl5esoft/lite-go/timeex"
	"github.com/go-redis/redis"
)

// INowTimeStartupContext is 当前时间启动上下文接口
type INowTimeStartupContext interface {
	GetRedisNowTimeOption() *redis.Options
}

// NewNowTimeStartupHandler is 当前时间启动处理器
func NewNowTimeStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(INowTimeStartupContext); ok {
			redis := New(
				sCtx.GetRedisNowTimeOption(),
			)
			ioc.Set(
				timeex.NowTimeIoCKey,
				redisex.NewNowTime(redis),
			)
		}

		return nil
	})
}
