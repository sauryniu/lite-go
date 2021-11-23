package mongosvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_factory_Db(t *testing.T) {
	t.Run("add - 非事务", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("add - 事务", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		uow := self.Uow()

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry, uow)
		err := db.Add(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.Len(t, res, 0)
	})

	t.Run("add - CreatedOn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer func() {
			ctrl.Finish()

			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		mockNowTime := contract.NewMockINowTime(ctrl)
		self.nowTime = mockNowTime

		unix := int64(11)
		mockNowTime.EXPECT().Unix().Return(unix)

		entry := testTimeAssociateModel{
			ID: "id",
		}
		db := self.Db(entry)
		err := db.Add(&entry)
		assert.NoError(t, err)

		var res []testTimeAssociateModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testTimeAssociateModel{
				{
					CreatedOn: unix,
					ID:        entry.ID,
				},
			},
		)
	})

	t.Run("remove - 非事务", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		err = db.Remove(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.Len(t, res, 0)
	})

	t.Run("remove - 事务", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		err = self.Db(
			entry,
			self.Uow(),
		).Remove(entry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("remove - DeletedOn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer func() {
			ctrl.Finish()

			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		mockNowTime := contract.NewMockINowTime(ctrl)
		self.nowTime = mockNowTime

		unix := int64(11)
		mockNowTime.EXPECT().Unix().Return(unix)

		entry := testTimeAssociateModel{
			ID: "id",
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		err = db.Remove(&entry)
		assert.NoError(t, err)

		var res []testTimeAssociateModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testTimeAssociateModel{
				{
					DeletedOn: unix,
					ID:        entry.ID,
				},
			},
		)
	})

	t.Run("save - 非事务", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		modifiedEntry := testModel{
			ID:   entry.ID,
			Name: "11",
			Age:  11,
		}
		err = db.Save(modifiedEntry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{modifiedEntry},
		)
	})

	t.Run("save - 非事务", func(t *testing.T) {
		defer func() {
			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		entry := testModel{
			ID:   "id",
			Name: "1",
			Age:  1,
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		modifiedEntry := testModel{
			ID:   entry.ID,
			Name: "11",
			Age:  11,
		}
		err = self.Db(
			modifiedEntry,
			self.Uow(),
		).Save(modifiedEntry)
		assert.NoError(t, err)

		var res []testModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testModel{entry},
		)
	})

	t.Run("save - ModifiedOn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer func() {
			ctrl.Finish()

			db, err := testClient.GetDb()
			assert.NoError(t, err)

			db.Drop(testClient.Ctx)
		}()

		self := new(factory)
		self.client = testClient

		mockNowTime := contract.NewMockINowTime(ctrl)
		self.nowTime = mockNowTime

		unix := int64(11)
		mockNowTime.EXPECT().Unix().Return(unix)

		entry := testTimeAssociateModel{
			ID: "id",
		}
		db := self.Db(entry)
		err := db.Add(entry)
		assert.NoError(t, err)

		err = db.Save(&entry)
		assert.NoError(t, err)

		var res []testTimeAssociateModel
		err = db.Query().ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(
			t,
			res,
			[]testTimeAssociateModel{
				{
					ID:         entry.ID,
					ModifiedOn: unix,
				},
			},
		)
	})
}
