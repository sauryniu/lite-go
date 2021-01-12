package db

import "github.com/ahl5esoft/lite-go/plugin/db/identity"

// IoCKey is 依赖注入键
const IoCKey = "db"

// IFactory is 数据工厂
type IFactory interface {
	Db(entry identity.IIdentity, extra ...interface{}) IRepository
	Uow() IUnitOfWork
}
