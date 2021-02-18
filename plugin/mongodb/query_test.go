package mongodb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_query_Count(t *testing.T) {
	res, err := newQuery(pool, testStruct).Count()
	assert.NoError(t, err)
	assert.Equal(
		t,
		res,
		int64(0),
	)
}

func Test_query_Count_Where(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	res, err := newQuery(pool, testStruct).Where(bson.M{
		"Age": 1,
	}).Count()
	assert.NoError(t, err)
	assert.Equal(
		t,
		res,
		int64(1),
	)
}

func Test_query_ToArray(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(&res)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry1, entry2},
	)
}

func Test_query_ToArray_Order(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).Order("Age").ToArray(&res)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry1, entry2},
	)
}

func Test_query_ToArray_OrderByDesc(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).OrderByDesc("Age").ToArray(&res)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry2, entry1},
	)
}

func Test_query_ToArray_Skip(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).Skip(1).ToArray(&res)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry2},
	)
}

func Test_query_ToArray_Take(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).Take(1).ToArray(&res)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry1},
	)
}

func Test_query_ToArray_dstIsReflectValue(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).ToArray(
		reflect.ValueOf(&res),
	)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry1, entry2},
	)
}

func Test_query_ToArray_Where(t *testing.T) {
	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	uow := newUnitOfWork(pool)
	entry1 := testModel{
		ID:   "id1",
		Name: "1",
		Age:  1,
	}
	uow.registerAdd(entry1)
	entry2 := testModel{
		ID:   "id2",
		Name: "2",
		Age:  2,
	}
	uow.registerAdd(entry2)
	err = uow.Commit()
	assert.NoError(t, err)

	var res []testModel
	err = newQuery(pool, testStruct).Where(bson.M{
		"_id": "id1",
	}).ToArray(&res)
	assert.NoError(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entry1},
	)
}
