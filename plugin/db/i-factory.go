package db

import "github.com/ahl5esoft/lite-go/plugin/db/identity"

// IFactory is 数据工厂
type IFactory interface {
	Db(entry identity.IIdentity, extra ...interface{}) IRepository
	Uow() IUnitOfWork
}
