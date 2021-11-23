package mongosvc

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_query_Count(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		res, err := newQuery(testClient, testModelMetadata).Count()
		assert.NoError(t, err)
		assert.Equal(
			t,
			res,
			int64(0),
		)
	})

	t.Run("where", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		res, err := newQuery(testClient, testModelMetadata).Where(bson.M{
			"Age": 1,
		}).Count()
		assert.NoError(t, err)
		assert.Equal(
			t,
			res,
			int64(1),
		)
	})
}

func Test_query_ToArray(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1, entry2},
		)
	})

	t.Run("dst is reflect.Value", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).ToArray(
			reflect.ValueOf(&res),
		)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1, entry2},
		)
	})

	t.Run("order", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).Order("Age").ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1, entry2},
		)
	})

	t.Run("orderByDesc", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).OrderByDesc("Age").ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry2, entry1},
		)
	})

	t.Run("skip", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).Skip(1).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry2},
		)
	})

	t.Run("take", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).Take(1).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1},
		)
	})

	t.Run("where", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		uow := newUnitOfWork(testClient)
		entry1 := testModel{
			ID:   "id1",
			Name: "1",
			Age:  1,
		}
		uow.RegisterAdd(entry1)
		entry2 := testModel{
			ID:   "id2",
			Name: "2",
			Age:  2,
		}
		uow.RegisterAdd(entry2)
		err := uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newQuery(testClient, testModelMetadata).Where(bson.M{
			"_id": "id1",
		}).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1},
		)
	})
}
