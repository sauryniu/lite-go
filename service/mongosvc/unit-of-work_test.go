package mongosvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_unitOfWork_registerAdd(t *testing.T) {
	t.Run("未提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "add",
			Age:  1,
		}
		uow := newUnitOfWork(testClient)
		uow.RegisterAdd(entry)
		assert.EqualValues(
			t,
			uow.addQueue,
			[]contract.IDbModel{entry},
		)
	})

	t.Run("提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "add",
			Age:  1,
		}
		uow := newUnitOfWork(testClient)
		uow.RegisterAdd(entry)
		err := uow.Commit()
		assert.NoError(t, err)
		assert.Len(t, uow.addQueue, 0)

		db, err := testClient.GetDb()
		assert.NoError(t, err)

		defer db.Drop(testClient.Ctx)

		name, err := testModelMetadata.GetTableName()
		assert.NoError(t, err)

		cur, err := db.Collection(name).Find(testClient.Ctx, bson.D{})
		assert.NoError(t, err)

		entries := make([]testModel, 0)
		for cur.Next(testClient.Ctx) {
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
	})
}

func Test_unitOfWork_registerRemove(t *testing.T) {
	t.Run("未提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "remove",
			Age:  1,
		}
		uow := newUnitOfWork(testClient)
		uow.RegisterRemove(entry)
		assert.EqualValues(
			t,
			uow.removeQueue,
			[]contract.IDbModel{entry},
		)
	})

	t.Run("提交", func(t *testing.T) {
		entry := testModel{
			ID:   "id",
			Name: "remove",
			Age:  1,
		}
		uow := newUnitOfWork(testClient)
		uow.RegisterAdd(entry)
		uow.RegisterRemove(entry)
		err := uow.Commit()
		assert.NoError(t, err)
		assert.Len(t, uow.removeQueue, 0)

		db, err := testClient.GetDb()
		assert.NoError(t, err)

		defer db.Drop(testClient.Ctx)

		name, err := testModelMetadata.GetTableName()
		assert.NoError(t, err)

		cur, err := db.Collection(name).Find(testClient.Ctx, bson.D{})
		assert.NoError(t, err)

		entries := make([]testModel, 0)
		for cur.Next(testClient.Ctx) {
			var temp testModel
			err = cur.Decode(&temp)
			assert.NoError(t, err)

			entries = append(entries, temp)
		}

		assert.Len(t, entries, 0)
	})
}

func Test_unitOfWork_registerSave(t *testing.T) {
	t.Run("未提交", func(t *testing.T) {
		uow := newUnitOfWork(testClient)
		entry := testModel{
			ID:   "id-2",
			Name: "save",
			Age:  2,
		}
		uow.RegisterSave(entry)
		assert.EqualValues(
			t,
			uow.saveQueue,
			[]contract.IDbModel{entry},
		)
	})

	t.Run("提交", func(t *testing.T) {
		uow := newUnitOfWork(testClient)
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
		assert.Len(t, uow.saveQueue, 0)

		db, err := testClient.GetDb()
		assert.NoError(t, err)

		defer db.Drop(testClient.Ctx)

		name, err := testModelMetadata.GetTableName()
		assert.NoError(t, err)

		cur, err := db.Collection(name).Find(testClient.Ctx, bson.D{})
		assert.NoError(t, err)

		entries := make([]testModel, 0)
		for cur.Next(testClient.Ctx) {
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
	})
}
