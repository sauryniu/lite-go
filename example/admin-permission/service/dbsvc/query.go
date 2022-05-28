package dbsvc

import (
	"github.com/ahl5esoft/lite-go/contract"
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"

	"go.mongodb.org/mongo-driver/bson"
)

type query struct {
	adminPermissionDbQuery contract.IDbQuery
	originDbQuery          contract.IDbQuery
	filter                 bson.M
	adminID                string
	model                  string
}

func (m *query) Count() (count int64, err error) {
	var filter bson.M
	if filter, err = m.getFilter(); err != nil {
		return
	}

	return m.originDbQuery.Where(filter).Count()
}

func (m *query) Order(fields ...string) contract.IDbQuery {
	m.originDbQuery.Order(fields...)
	return m
}

func (m *query) OrderByDesc(fields ...string) contract.IDbQuery {
	m.originDbQuery.OrderByDesc(fields...)
	return m
}

func (m *query) Skip(v int) contract.IDbQuery {
	m.originDbQuery.Skip(v)
	return m
}

func (m *query) Take(v int) contract.IDbQuery {
	m.originDbQuery.Take(v)
	return m
}

func (m *query) ToArray(dst interface{}) (err error) {
	var filter bson.M
	if filter, err = m.getFilter(); err != nil {
		return
	}

	err = m.originDbQuery.Where(filter).ToArray(dst)
	return
}

func (m *query) Where(args ...interface{}) contract.IDbQuery {
	if len(args) == 0 {
		return m
	}

	if f, ok := args[0].(bson.M); ok {
		m.filter = f
	}
	return m
}

func (m *query) getFilter() (filter bson.M, err error) {
	filter = m.filter
	m.filter = nil

	if filter == nil {
		filter = bson.M{}
	}

	var adminPermissions []global.AdminPermission
	err = m.adminPermissionDbQuery.Where(bson.M{
		"_id": m.adminID,
	}).ToArray(&adminPermissions)
	if err != nil {
		return
	}

	if len(adminPermissions) > 0 {
		if v, ok := adminPermissions[0].Permission[m.model]; ok {
			if cv, ok := v[dbop.Query]; ok {
				for sk, sv := range cv.(bson.M) {
					filter[sk] = sv
				}
			}
		}
	}
	return
}
