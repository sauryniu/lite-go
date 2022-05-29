package dbsvc

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"
	dbsvc_ "github.com/ahl5esoft/lite-go/service/dbsvc"

	"go.mongodb.org/mongo-driver/bson"
)

type dbFactory struct {
	mongoDbFactory   contract.IDbFactory
	adminID          string
	adminPermissions []global.AdminPermission
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

	return dbsvc_.NewDbRepository(uow, isTx, func() contract.IDbQuery {
		modelType := reflect.TypeOf(entry)
		return newDbQuery(
			m.mongoDbFactory.Db(entry).Query(),
			modelType.Name(),
			m,
		)
	})
}

func (m *dbFactory) GetAdminPermissions() ([]global.AdminPermission, error) {
	if m.adminPermissions == nil {
		var adminPermissions []global.AdminPermission
		err := m.mongoDbFactory.Db(global.AdminPermission{}).Query().Where(bson.M{
			"_id": m.adminID,
		}).ToArray(&adminPermissions)
		if err != nil {
			return nil, err
		}

		m.adminPermissions = adminPermissions
	}

	return m.adminPermissions, nil
}

func (m *dbFactory) Uow() contract.IUnitOfWork {
	uow := m.mongoDbFactory.Uow()
	return newUnitOfWorkRepository(
		uow.(contract.IUnitOfWorkRepository),
		m,
	)
}

// 创建数据库工厂
func NewDbFactory(
	mongoDbFactory contract.IDbFactory,
	adminID string,
) contract.IDbFactory {
	return &dbFactory{
		mongoDbFactory: mongoDbFactory,
		adminID:        adminID,
	}
}
