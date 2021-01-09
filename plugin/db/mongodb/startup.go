package mongodb

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
)

// StartupContext is 启动上下文
type StartupContext struct {
	MongoOption FactoryOption
}

// NewStartup is 启动处理器
func NewStartup() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(*StartupContext); ok {
			f, err := New(sCtx.MongoOption)
			if err != nil {
				return err
			}

			ioc.Set("db", f)
		}
		return nil
	})
}
