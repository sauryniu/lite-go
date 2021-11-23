package mongosvc

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/service/dbsvc"
)

type factory struct {
	nowTime contract.INowTime
	client  *clientWrapper
}

func (m *factory) Db(entry contract.IDbModel, extra ...interface{}) contract.IDbRepository {
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

	model := getModelMetadata(
		reflect.TypeOf(entry),
	)
	return dbsvc.NewRepositoryBase(func() contract.IDbQuery {
		return newQuery(m.client, model)
	}, isTx, m.nowTime, uow)
}

func (m *factory) Uow() contract.IUnitOfWork {
	return newUnitOfWork(m.client)
}

func New(
	nowTime contract.INowTime,
	name string,
	uri string,
) contract.IDbFactory {
	return &factory{
		client:  newClientWrapper(name, uri),
		nowTime: nowTime,
	}
}
