package mysqldb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/jmoiron/sqlx"
)

type query struct {
	DB        *sqlx.DB
	ModelType reflect.Type
	Option    queryOption
}

func (m *query) Count() (count int64, err error) {
	var sql string
	if sql, err = newSQLMaker(m.ModelType).GetCount(m.Option); err != nil {
		return
	}

	err = m.DB.Get(&count, sql, m.Option.WhereArgs...)
	return
}

func (m *query) Order(fields ...string) db.IQuery {
	return m.order("ASC", fields...)
}

func (m *query) OrderByDesc(fields ...string) db.IQuery {
	return m.order("DESC", fields...)
}

func (m *query) Skip(v int) db.IQuery {
	m.Option.Skip = v
	return m
}

func (m *query) Take(v int) db.IQuery {
	m.Option.Take = v
	return m
}

func (m *query) ToArray(dst interface{}) (err error) {
	var sql string
	if sql, err = newSQLMaker(m.ModelType).GetSelect(m.Option); err != nil {
		return
	}

	err = m.DB.Select(dst, sql, m.Option.WhereArgs...)
	return
}

func (m *query) Where(args ...interface{}) db.IQuery {
	underscore.Chain(args).Each(func(r interface{}, i int) {
		if i == 0 {
			m.Option.Where = r.(string)
			m.Option.WhereArgs = make([]interface{}, 0)
		} else {
			m.Option.WhereArgs = append(m.Option.WhereArgs, r)
		}
	})
	return m
}

func (m *query) order(flag string, fields ...string) db.IQuery {
	underscore.Chain(fields).Each(func(r string, _ int) {
		m.Option.Orders = append(m.Option.Orders, orderOption{
			Field: r,
			Flag:  flag,
		})
	})
	return m
}

func newQuery(sqlxDB *sqlx.DB, modelType reflect.Type) *query {
	return &query{
		DB:        sqlxDB,
		ModelType: modelType,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
}
