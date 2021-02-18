package mysqldb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unitOfWork_Commit(t *testing.T) {
	assert.NotNil(t, sqlxDB)
	self := newUnitOfWork(sqlxDB)

	entry := testModel{
		ID:   "id-1",
		Name: "add",
		Age:  11,
	}
	self.RegisterAdd(entry)
	entry.Name = "save"
	self.RegisterSave(entry)
	self.RegisterRemove(entry)
	err := self.Commit()
	assert.Nil(t, err)
}

func Test_unitOfWork_RegisterAdd(t *testing.T) {
	self := newUnitOfWork(sqlxDB)

	entry := testModel{
		ID:   "id-1",
		Name: "add",
		Age:  11,
	}
	self.RegisterAdd(entry)

	sql, _ := newSQLMaker(
		reflect.TypeOf(entry),
	).GetAdd()
	assert.EqualValues(t, self.Items, []unitOfWorkItem{
		{
			Args: []interface{}{entry.ID, entry.Name, entry.Age},
			SQL:  sql,
		},
	})
}

func Test_unitOfWork_RegisterRemove(t *testing.T) {
	self := newUnitOfWork(sqlxDB)

	entry := testModel{
		ID:   "id-2",
		Name: "remove",
		Age:  11,
	}
	self.RegisterRemove(entry)

	sql, _ := newSQLMaker(
		reflect.TypeOf(entry),
	).GetRemove()
	assert.EqualValues(t, self.Items, []unitOfWorkItem{
		{
			Args: []interface{}{entry.ID},
			SQL:  sql,
		},
	})
}

func Test_unitOfWork_RegisterSave(t *testing.T) {
	self := newUnitOfWork(sqlxDB)

	entry := testModel{
		ID:   "id-3",
		Name: "save",
		Age:  11,
	}
	self.RegisterSave(entry)

	sql, _ := newSQLMaker(
		reflect.TypeOf(entry),
	).GetSave()
	assert.EqualValues(t, self.Items, []unitOfWorkItem{
		{
			Args: []interface{}{entry.Name, entry.Age, entry.ID},
			SQL:  sql,
		},
	})
}
