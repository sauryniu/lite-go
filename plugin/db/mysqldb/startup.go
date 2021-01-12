package mysqldb

import (
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/db"
)

// IStartupContext is 启动上下文接口
type IStartupContext interface {
	GetMysqlOption() FactoryOption
}

// NewStartupHandler is 启动处理器
func NewStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(IStartupContext); ok {
			f, err := New(
				sCtx.GetMysqlOption(),
			)
			if err != nil {
				return err
			}

			ioc.Set(db.IoCKey, f)
		}
		return nil
	})
}
