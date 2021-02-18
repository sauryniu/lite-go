package mysqldb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repository_Add(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "add",
		Age:  1,
	}
	self := repository{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entry),
		Uow: &unitOfWork{
			DB:    sqlxDB,
			Items: make([]unitOfWorkItem, 0),
		},
	}
	err := self.Add(entry)

	var res []testModel
	self.Query().ToArray(&res)

	self.Uow.RegisterRemove(entry)
	self.Uow.Commit()

	assert.Nil(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry},
	)
}

func Test_repository_Add_有事务(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "add-tx",
		Age:  1,
	}
	self := repository{
		DB:        sqlxDB,
		IsTx:      true,
		ModelType: reflect.TypeOf(entry),
		Uow: &unitOfWork{
			DB:    sqlxDB,
			Items: make([]unitOfWorkItem, 0),
		},
	}
	err := self.Add(entry)

	var res []testModel
	self.Query().ToArray(&res)

	assert.Nil(t, err)
	assert.Empty(t, res)
}

func Test_repository_Remove(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "remove",
		Age:  1,
	}
	self := repository{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entry),
		Uow: &unitOfWork{
			DB:    sqlxDB,
			Items: make([]unitOfWorkItem, 0),
		},
	}

	self.Uow.RegisterAdd(entry)
	self.Uow.Commit()

	err := self.Remove(entry)

	var res []testModel
	self.Query().ToArray(&res)

	assert.Nil(t, err)
	assert.Empty(t, res)
}

func Test_repository_Remove_事务(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "remove-tx",
		Age:  1,
	}
	self := repository{
		DB:        sqlxDB,
		IsTx:      true,
		ModelType: reflect.TypeOf(entry),
		Uow: &unitOfWork{
			DB:    sqlxDB,
			Items: make([]unitOfWorkItem, 0),
		},
	}

	self.Uow.RegisterAdd(entry)
	self.Uow.Commit()

	err := self.Remove(entry)

	var res []testModel
	self.Query().ToArray(&res)

	self.Uow.Commit()

	assert.Nil(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry},
	)
}

func Test_repository_Save(t *testing.T) {
	entry := testModel{
		ID:   "id-save",
		Name: "add",
		Age:  1,
	}
	self := repository{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entry),
		Uow: &unitOfWork{
			DB:    sqlxDB,
			Items: make([]unitOfWorkItem, 0),
		},
	}

	self.Uow.RegisterAdd(entry)
	self.Uow.Commit()

	entry.Name = "save"
	err := self.Save(entry)

	var res []testModel
	self.Query().ToArray(&res)

	self.Uow.RegisterRemove(entry)
	self.Uow.Commit()

	assert.Nil(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0].Name, "save")
}

func Test_repository_Save_事务(t *testing.T) {
	entry := testModel{
		ID:   "id-save",
		Name: "add",
		Age:  1,
	}
	self := repository{
		DB:        sqlxDB,
		IsTx:      true,
		ModelType: reflect.TypeOf(entry),
		Uow: &unitOfWork{
			DB:    sqlxDB,
			Items: make([]unitOfWorkItem, 0),
		},
	}

	self.Uow.RegisterAdd(entry)
	self.Uow.Commit()

	entry.Name = "save"
	err := self.Save(entry)

	var res []testModel
	self.Query().ToArray(&res)

	self.Uow.RegisterRemove(entry)
	self.Uow.Commit()

	assert.Nil(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0].Name, "add")
}
