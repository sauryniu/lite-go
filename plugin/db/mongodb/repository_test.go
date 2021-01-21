package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repository_Add(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	err = newRepository(pool, testStruct, uow, false).Add(entry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry},
	)
}

func Test_repository_Add_WithTx(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	err = newRepository(pool, testStruct, uow, true).Add(entry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.Len(t, res, 0)
}

func Test_repository_Query(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.NoError(t, err)
	assert.Len(t, res, 0)
}

func Test_repository_Remove(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry)
	err = uow.Commit()
	assert.NoError(t, err)

	err = newRepository(pool, testStruct, uow, false).Remove(entry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.Len(t, res, 0)
}

func Test_repository_Remove_WithTx(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry)
	err = uow.Commit()
	assert.NoError(t, err)

	err = newRepository(pool, testStruct, uow, true).Remove(entry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry},
	)
}

func Test_repository_Save(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry)
	err = uow.Commit()
	assert.NoError(t, err)

	modifiedEntry := testModel{
		ID:   entry.ID,
		Name: "11",
		Age:  11,
	}
	err = newRepository(pool, testStruct, uow, false).Save(modifiedEntry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{modifiedEntry},
	)
}

func Test_repository_Save_WithTx(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry)
	err = uow.Commit()
	assert.NoError(t, err)

	modifiedEntry := testModel{
		ID:   entry.ID,
		Name: "11",
		Age:  11,
	}
	err = newRepository(pool, testStruct, uow, true).Save(modifiedEntry)
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry},
	)
}
