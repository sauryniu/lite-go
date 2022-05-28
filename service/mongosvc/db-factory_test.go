package mongosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factory_Db(t *testing.T) {
	t.Run("delete - 非事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		self := new(dbFactory)
		self.driverFactory = testDriverFactory

		entry := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		err = self.Db(entry).Insert(entry)
		assert.NoError(t, err)

		err = self.Db(entry).Delete(entry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Db(entry).Query().ToArray(&res)
		assert.NoError(t, err)
		assert.Len(t, res, 0)
	})

	t.Run("delete - 事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		self := new(dbFactory)
		self.driverFactory = testDriverFactory

		entry := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		err = self.Db(entry).Insert(entry)
		assert.NoError(t, err)

		uow := self.Uow()
		err = self.Db(entry, uow).Delete(entry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Db(entry).Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("insert - 非事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		self := new(dbFactory)
		self.driverFactory = testDriverFactory

		entry := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		err = self.Db(entry).Insert(entry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Db(entry).Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("insert - 事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		self := new(dbFactory)
		self.driverFactory = testDriverFactory

		uow := self.Uow()
		entry := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		err = self.Db(entry, uow).Insert(entry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Db(entry).Query().ToArray(&res)
		assert.NoError(t, err)
		assert.Len(t, res, 0)
	})

	t.Run("update - 非事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		self := new(dbFactory)
		self.driverFactory = testDriverFactory

		insertEntry := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		err = self.Db(insertEntry).Insert(insertEntry)
		assert.NoError(t, err)

		updateEntry := testModel{
			ID:  insertEntry.ID,
			Int: insertEntry.Int + 1,
			Str: insertEntry.Str,
		}
		err = self.Db(insertEntry).Update(updateEntry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Db(insertEntry).Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{updateEntry},
		)
	})

	t.Run("update - 事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		self := new(dbFactory)
		self.driverFactory = testDriverFactory

		insertEntry := testModel{
			ID:  "id1",
			Int: 1,
			Str: "2",
		}
		db := self.Db(insertEntry)
		err = db.Insert(insertEntry)
		assert.NoError(t, err)

		uow := self.Uow()
		updateEntry := testModel{
			ID:  insertEntry.ID,
			Int: insertEntry.Int + 1,
			Str: insertEntry.Str,
		}
		err = self.Db(updateEntry, uow).Update(updateEntry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Db(insertEntry).Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{insertEntry},
		)
	})
}
