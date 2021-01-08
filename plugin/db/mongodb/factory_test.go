package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factory_Db(t *testing.T) {
	f, err := New(FactoryOption{
		DbName: "lite-go",
		URI:    "mongodb://localhost:27017",
	})
	assert.NoError(t, err)

	var res []testModel
	err = f.Db(testModel{}).Query().ToArray(&res)
	assert.NoError(t, err)
	assert.Len(t, res, 0)
}

func Test_factory_Uow(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.Ctx)

	f, err := New(FactoryOption{
		DbName: "lite-go",
		URI:    "mongodb://localhost:27017",
	})
	assert.NoError(t, err)

	uow := f.Uow()
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	err = f.Db(testModel{}, uow).Add(entry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.Len(t, res, 0)
}

func Test_factory_Uow_Commit(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.Ctx)

	f, err := New(FactoryOption{
		DbName: "lite-go",
		URI:    "mongodb://localhost:27017",
	})
	assert.NoError(t, err)

	uow := f.Uow()
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	err = f.Db(testModel{}, uow).Add(entry)
	assert.NoError(t, err)

	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry},
	)
}
