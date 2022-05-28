package mongosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_query_Count(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		uow.RegisterInsert(testModel{
			ID:  "id1",
			Int: 1,
			Str: "1",
		})
		uow.RegisterInsert(testModel{
			ID:  "id2",
			Int: 1,
			Str: "2",
		})
		err = uow.Commit()
		assert.NoError(t, err)

		res, err := newDbQuery(testDriverFactory, testModelMetadata).Count()
		assert.NoError(t, err)
		assert.Equal(
			t,
			res,
			int64(2),
		)
	})

	t.Run("where", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		uow.RegisterInsert(testModel{
			ID:  "id1",
			Int: 1,
			Str: "1",
		})
		uow.RegisterInsert(testModel{
			ID:  "id2",
			Int: 1,
			Str: "2",
		})
		err = uow.Commit()
		assert.NoError(t, err)

		res, err := newDbQuery(testDriverFactory, testModelMetadata).Where(bson.M{
			"str": "2",
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
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry1 := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		uow.RegisterInsert(entry1)
		entry2 := testModel{
			ID:  "id2",
			Int: 2,
			Str: "1",
		}
		uow.RegisterInsert(entry2)
		err = uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newDbQuery(testDriverFactory, testModelMetadata).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1, entry2},
		)
	})

	t.Run("order", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry1 := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		uow.RegisterInsert(entry1)
		entry2 := testModel{
			ID:  "id2",
			Int: 2,
			Str: "1",
		}
		uow.RegisterInsert(entry2)
		err = uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newDbQuery(testDriverFactory, testModelMetadata).Order("str").ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry2, entry1},
		)
	})

	t.Run("orderByDesc", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry1 := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		uow.RegisterInsert(entry1)
		entry2 := testModel{
			ID:  "id2",
			Int: 2,
			Str: "1",
		}
		uow.RegisterInsert(entry2)
		err = uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newDbQuery(testDriverFactory, testModelMetadata).OrderByDesc("Int").ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry2, entry1},
		)
	})

	t.Run("skip", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry1 := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		uow.RegisterInsert(entry1)
		entry2 := testModel{
			ID:  "id2",
			Int: 2,
			Str: "1",
		}
		uow.RegisterInsert(entry2)
		err = uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newDbQuery(testDriverFactory, testModelMetadata).Skip(1).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry2},
		)
	})

	t.Run("take", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry1 := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		uow.RegisterInsert(entry1)
		entry2 := testModel{
			ID:  "id2",
			Int: 2,
			Str: "1",
		}
		uow.RegisterInsert(entry2)
		err = uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newDbQuery(testDriverFactory, testModelMetadata).Take(1).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1},
		)
	})

	t.Run("where", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry1 := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		uow.RegisterInsert(entry1)
		entry2 := testModel{
			ID:  "id2",
			Int: 2,
			Str: "1",
		}
		uow.RegisterInsert(entry2)
		err = uow.Commit()
		assert.NoError(t, err)

		var res []testModel
		err = newDbQuery(testDriverFactory, testModelMetadata).Where(bson.M{
			"_id": entry1.ID,
		}).ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry1},
		)
	})
}
