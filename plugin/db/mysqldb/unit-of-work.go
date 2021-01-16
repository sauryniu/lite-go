package mysqldb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"github.com/jmoiron/sqlx"
)

type unitOfWorkItem struct {
	Args []interface{}
	SQL  string
}

type unitOfWork struct {
	DB    *sqlx.DB
	Items []unitOfWorkItem
}

func (m *unitOfWork) Commit() error {
	tx := m.DB.MustBegin()
	underscore.Chain(m.Items).Each(func(r unitOfWorkItem, _ int) {
		tx.MustExec(r.SQL, r.Args...)
	})
	m.Items = make([]unitOfWorkItem, 0)
	return tx.Commit()
}

func (m *unitOfWork) RegisterAdd(entry identity.IIdentity) (err error) {
	modelType := reflect.TypeOf(entry)
	modelValue := reflect.ValueOf(entry)

	var args []interface{}
	underscore.Chain(
		identity.NewStruct(modelType).FindFields(),
	).Map(func(r identity.IField, _ int) interface{} {
		return r.GetValue(modelValue)
	}).Value(&args)

	var sql string
	if sql, err = newSQLMaker(modelType).GetAdd(); err != nil {
		return
	}

	m.Items = append(m.Items, unitOfWorkItem{
		Args: args,
		SQL:  sql,
	})
	return
}

func (m *unitOfWork) RegisterRemove(entry identity.IIdentity) (err error) {
	modelType := reflect.TypeOf(entry)

	var column identity.IField
	if column, err = identity.NewStruct(modelType).GetIDField(); err != nil {
		return
	}

	sql, _ := newSQLMaker(modelType).GetRemove()
	m.Items = append(m.Items, unitOfWorkItem{
		Args: []interface{}{
			column.GetValue(
				reflect.ValueOf(entry),
			),
		},
		SQL: sql,
	})
	return
}

func (m *unitOfWork) RegisterSave(entry identity.IIdentity) (err error) {
	modelType := reflect.TypeOf(entry)
	entryValue := reflect.ValueOf(entry)

	var args []interface{}
	var pk interface{}
	underscore.Chain(
		identity.NewStruct(modelType).FindFields(),
	).Where(func(r identity.IField, _ int) bool {
		ok := r.GetStructName() != ""
		if ok {
			pk = r.GetValue(entryValue)
		}
		return !ok
	}).Map(func(r identity.IField, _ int) interface{} {
		return r.GetValue(entryValue)
	}).Value(&args)
	args = append(args, pk)

	var sql string
	if sql, err = newSQLMaker(modelType).GetSave(); err != nil {
		return
	}

	m.Items = append(m.Items, unitOfWorkItem{
		Args: args,
		SQL:  sql,
	})
	return
}

func newUnitOfWork(sqlxDB *sqlx.DB) *unitOfWork {
	return &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
}
