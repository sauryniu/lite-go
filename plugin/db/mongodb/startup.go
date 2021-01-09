package mongodb

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
)

// IStartupContext is 启动上下文接口
type IStartupContext interface {
	GetMongoOption() FactoryOption
}

// NewStartup is 启动处理器
func NewStartup() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(IStartupContext); ok {
			f, err := New(
				sCtx.GetMongoOption(),
			)
			if err != nil {
				return err
			}

			ioc.Set("db", f)
		}
		return nil
	})
}
