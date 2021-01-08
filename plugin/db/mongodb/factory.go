package mongodb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
)

// FactoryOption is 工厂选项
type FactoryOption struct {
	DbName string
	URI    string
}

type factory struct {
	Pool *connectPool
}

func (m factory) Db(entry identity.IIdentity, extra ...interface{}) db.IRepository {
	var uow *unitOfWork
	isTx := true
	underscore.Chain(extra).Each(func(r interface{}, _ int) {
		if v, ok := r.(*unitOfWork); ok {
			uow = v
		}
	})

	if uow == nil {
		isTx = false
		uow = m.Uow().(*unitOfWork)
	}

	s := identity.NewStruct(
		reflect.TypeOf(entry),
	)
	return newRepository(m.Pool, s, uow, isTx)
}

func (m factory) Uow() db.IUnitOfWork {
	return newUnitOfWork(m.Pool)
}

// New is 创建db.IDbFactory
func New(opt FactoryOption) (db.IFactory, error) {
	pool := newPool(opt.DbName, opt.URI)
	if _, err := pool.GetClient(); err != nil {
		return nil, err
	}

	return &factory{
		Pool: pool,
	}, nil
}
