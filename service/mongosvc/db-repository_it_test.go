package mongosvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/service/dbsvc"

	"github.com/stretchr/testify/assert"
)

func Test_Repository_Delete(t *testing.T) {
	t.Run("无事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry := testModel{
			ID:  "id-1",
			Int: 1,
			Str: "add",
		}
		uow.RegisterInsert(entry)
		err = uow.Commit()
		assert.NoError(t, err)

		self := dbsvc.NewRepository(uow, false, func() contract.IDbQuery {
			return newDbQuery(testDriverFactory, model)
		})
		err = self.Delete(entry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Query().ToArray(&res)
		assert.NoError(t, err)

		assert.Len(t, res, 0)
	})
}

func Test_Repository_Insert(t *testing.T) {
	t.Run("无事务", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)
		entry := testModel{
			ID:  "id-1",
			Int: 1,
			Str: "add",
		}
		self := dbsvc.NewRepository(uow, false, func() contract.IDbQuery {
			return newDbQuery(testDriverFactory, model)
		})
		err = self.Insert(entry)
		assert.NoError(t, err)

		var res []testModel
		err = self.Query().ToArray(&res)
		assert.NoError(t, err)

		assert.EqualValues(t, res, []testModel{entry})
	})
}
