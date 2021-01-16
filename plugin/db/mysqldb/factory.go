package mysqldb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"github.com/jmoiron/sqlx"
)

type factory struct {
	DB *sqlx.DB
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

	return repository{
		DB:        m.DB,
		IsTx:      isTx,
		ModelType: reflect.TypeOf(entry),
		Uow:       uow,
	}
}

func (m factory) Uow() db.IUnitOfWork {
	return newUnitOfWork(m.DB)
}

// New is 创建db.IDbFactory
func New(connString string) (f db.IFactory, err error) {
	var db *sqlx.DB
	if db, err = sqlx.Open("mysql", connString); err != nil {
		return
	}

	f = &factory{
		DB: db,
	}
	return
}
