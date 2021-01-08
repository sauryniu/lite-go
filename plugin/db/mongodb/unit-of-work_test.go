package mongodb

import (
	"testing"

	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_unitOfWork_RegisterAdd(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "add",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.RegisterAdd(entry)
	assert.EqualValues(
		t,
		uow.AddQueue,
		[]identity.IIdentity{entry},
	)
}

func Test_unitOfWork_RegisterAdd_WithCommit(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "add",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.RegisterAdd(entry)
	err := uow.Commit()
	assert.NoError(t, err)
	assert.Len(t, uow.AddQueue, 0)

	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.Ctx)

	name, err := testStruct.GetName()
	assert.NoError(t, err)

	cur, err := db.Collection(name).Find(pool.Ctx, bson.D{})
	assert.NoError(t, err)

	entries := make([]testModel, 0)
	for cur.Next(pool.Ctx) {
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

func Test_unitOfWork_RegisterRemove(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "remove",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.RegisterRemove(entry)
	assert.EqualValues(
		t,
		uow.RemoveQueue,
		[]identity.IIdentity{entry},
	)
}

func Test_unitOfWork_RegisterRemove_WithCommit(t *testing.T) {
	entry := testModel{
		ID:   "id",
		Name: "remove",
		Age:  1,
	}
	uow := newUnitOfWork(pool)
	uow.RegisterAdd(entry)
	uow.RegisterRemove(entry)
	err := uow.Commit()
	assert.NoError(t, err)
	assert.Len(t, uow.RemoveQueue, 0)

	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.Ctx)

	name, err := testStruct.GetName()
	assert.NoError(t, err)

	cur, err := db.Collection(name).Find(pool.Ctx, bson.D{})
	assert.NoError(t, err)

	entries := make([]testModel, 0)
	for cur.Next(pool.Ctx) {
		var temp testModel
		err = cur.Decode(&temp)
		assert.NoError(t, err)

		entries = append(entries, temp)
	}

	assert.Len(t, entries, 0)
}

func Test_unitOfWork_RegisterSave(t *testing.T) {
	uow := newUnitOfWork(pool)
	entry := testModel{
		ID:   "id-2",
		Name: "save",
		Age:  2,
	}
	uow.RegisterSave(entry)
	assert.EqualValues(
		t,
		uow.SaveQueue,
		[]identity.IIdentity{entry},
	)
}

func Test_unitOfWork_RegisterSave_WithCommit(t *testing.T) {
	uow := newUnitOfWork(pool)
	uow.RegisterAdd(testModel{
		ID:   "id-2",
		Name: "add",
		Age:  1,
	})
	entry := testModel{
		ID:   "id-2",
		Name: "save",
		Age:  2,
	}
	uow.RegisterSave(entry)
	err := uow.Commit()
	assert.NoError(t, err)
	assert.Len(t, uow.SaveQueue, 0)

	db, err := pool.GetDb()
	assert.NoError(t, err)

	defer db.Drop(pool.Ctx)

	name, err := testStruct.GetName()
	assert.NoError(t, err)

	cur, err := db.Collection(name).Find(pool.Ctx, bson.D{})
	assert.NoError(t, err)

	entries := make([]testModel, 0)
	for cur.Next(pool.Ctx) {
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
