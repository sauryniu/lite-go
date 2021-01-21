package mongodb

import (
	"testing"

	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_unitOfWork_registerAdd(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "add",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.registerAdd(entry)
	assert.EqualValues(
		t,
		uow.addQueue,
		[]identity.IIdentity{entry},
	)
}

func Test_unitOfWork_registerAdd_WithCommit(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "add",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.registerAdd(entry)
	err := uow.Commit()
	assert.NoError(t, err)
	assert.Len(t, uow.addQueue, 0)

	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	name, err := testStruct.GetName()
	assert.NoError(t, err)

	cur, err := db.Collection(name).Find(pool.ctx, bson.D{})
	assert.NoError(t, err)

	entries := make([]testModel, 0)
	for cur.Next(pool.ctx) {
		var temp testModel
		err = cur.Decode(&temp)
		assert.NoError(t, err)

		entries = append(entries, temp)
	}

	assert.EqualValues(
		t,
		entries,
		[]testModel{entry},
	)
}

func Test_unitOfWork_registerRemove(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "remove",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.registerRemove(entry)
	assert.EqualValues(
		t,
		uow.removeQueue,
		[]identity.IIdentity{entry},
	)
}

func Test_unitOfWork_registerRemove_WithCommit(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "remove",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.registerAdd(entry)
	uow.registerRemove(entry)
	err := uow.Commit()
	assert.NoError(t, err)
	assert.Len(t, uow.removeQueue, 0)

	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	name, err := testStruct.GetName()
	assert.NoError(t, err)

	cur, err := db.Collection(name).Find(pool.ctx, bson.D{})
	assert.NoError(t, err)

	entries := make([]testModel, 0)
	for cur.Next(pool.ctx) {
		var temp testModel
		err = cur.Decode(&temp)
		assert.NoError(t, err)

		entries = append(entries, temp)
	}

	assert.Len(t, entries, 0)
}

func Test_unitOfWork_registerSave(t *testing.T) {
	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id-2",
		Name: "save",
		Age:  2,
	}
	uow.registerSave(entry)
	assert.EqualValues(
		t,
		uow.saveQueue,
		[]identity.IIdentity{entry},
	)
}

func Test_unitOfWork_registerSave_WithCommit(t *testing.T) {
	uow := newUnitOfWork(pool)
	uow.registerAdd(testModel{
		ID:   "id-2",
		Name: "add",
		Age:  1,
	})
	entry := testModel{
		ID:   "id-2",
		Name: "save",
		Age:  2,
	}
	uow.registerSave(entry)
	err := uow.Commit()
	assert.NoError(t, err)
	assert.Len(t, uow.saveQueue, 0)

	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.ctx)

	name, err := testStruct.GetName()
	assert.NoError(t, err)

	cur, err := db.Collection(name).Find(pool.ctx, bson.D{})
	assert.NoError(t, err)

	entries := make([]testModel, 0)
	for cur.Next(pool.ctx) {
		var temp testModel
		err = cur.Decode(&temp)
		assert.NoError(t, err)

		entries = append(entries, temp)
	}

	assert.EqualValues(
		t,
		entries,
		[]testModel{entry},
	)
}
