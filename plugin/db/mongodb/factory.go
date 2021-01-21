package mongodb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
)

type factory struct {
	pool *connectPool
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

	modelStruct := identity.NewStruct(
		reflect.TypeOf(entry),
	)
	return newRepository(m.pool, modelStruct, uow, isTx)
}

func (m factory) Uow() db.IUnitOfWork {
	return newUnitOfWork(m.pool)
}

// New is 创建db.IDbFactory
func New(uri, dbName string) (db.IFactory, error) {
	pool := newPool(uri, dbName)
	if _, err := pool.GetClient(); err != nil {
		return nil, err
	}

	return &factory{
		pool: pool,
	}, nil
}
