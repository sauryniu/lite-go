package mongosvc

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/service/dbsvc"
)

type dbFactory struct {
	driverFactory *driverFactory
}

func (m *dbFactory) Db(entry contract.IDbModel, extra ...interface{}) contract.IDbRepository {
	var uow *unitOfWorkRepository
	isTx := true
	underscore.Chain(extra).Each(func(r interface{}, _ int) {
		if v, ok := r.(*unitOfWorkRepository); ok {
			uow = v
		}
	})

	if uow == nil {
		isTx = false
		uow = m.Uow().(*unitOfWorkRepository)
	}

	model := getModelMetadata(
		reflect.TypeOf(entry),
	)
	return dbsvc.NewRepository(uow, isTx, func() contract.IDbQuery {
		return newDbQuery(m.driverFactory, model)
	})
}

func (m *dbFactory) Uow() contract.IUnitOfWork {
	return newUnitOfWorkRepository(m.driverFactory)
}

// 创建数据库工厂
func NewDbFactory(name string, uri string) contract.IDbFactory {
	return &dbFactory{
		driverFactory: newDriverFactory(name, uri),
	}
}
