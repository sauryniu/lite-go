package dbsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbQuery struct {
	originDbQuery contract.IDbQuery
	filter        bson.M
	model         string
	dbFactory     *dbFactory
}

func (m *dbQuery) Count() (count int64, err error) {
	var filter bson.M
	if filter, err = m.getFilter(); err != nil {
		return
	}

	return m.originDbQuery.Where(filter).Count()
}

func (m *dbQuery) Order(fields ...string) contract.IDbQuery {
	m.originDbQuery.Order(fields...)
	return m
}

func (m *dbQuery) OrderByDesc(fields ...string) contract.IDbQuery {
	m.originDbQuery.OrderByDesc(fields...)
	return m
}

func (m *dbQuery) Skip(v int) contract.IDbQuery {
	m.originDbQuery.Skip(v)
	return m
}

func (m *dbQuery) Take(v int) contract.IDbQuery {
	m.originDbQuery.Take(v)
	return m
}

func (m *dbQuery) ToArray(dst interface{}) (err error) {
	var filter bson.M
	if filter, err = m.getFilter(); err != nil {
		return
	}

	err = m.originDbQuery.Where(filter).ToArray(dst)
	return
}

func (m *dbQuery) Where(args ...interface{}) contract.IDbQuery {
	if len(args) == 0 {
		return m
	}

	if f, ok := args[0].(bson.M); ok {
		m.filter = f
	}
	return m
}

func (m *dbQuery) getFilter() (filter bson.M, err error) {
	filter = m.filter
	m.filter = nil

	if filter == nil {
		filter = bson.M{}
	}

	var adminPermissions []global.AdminPermission
	if adminPermissions, err = m.dbFactory.GetAdminPermissions(); err != nil {
		return
	}

	if len(adminPermissions) > 0 {
		if v, ok := adminPermissions[0].Permission[m.model]; ok {
			if cv, ok := v[dbop.Query]; ok {
				var doc bson.M
				bson.Unmarshal(
					cv.(primitive.Binary).Data,
					&doc,
				)
				for sk, sv := range doc {
					filter[sk] = sv
				}
			}
		}
	}
	return
}

func newDbQuery(
	originDbQuery contract.IDbQuery,
	model string,
	dbFactory *dbFactory,
) contract.IDbQuery {
	return &dbQuery{
		dbFactory:     dbFactory,
		originDbQuery: originDbQuery,
		model:         model,
	}
}
