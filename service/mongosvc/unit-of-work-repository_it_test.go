package mongosvc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_unitOfWorkRepository_Commit(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		model := getModelMetadata(testModelType)
		c, err := testDriverFactory.BuildCollection(model)
		assert.NoError(t, err)

		defer c.Drop(testDriverFactory.Ctx)

		uow := newUnitOfWorkRepository(testDriverFactory)

		uow.RegisterInsert(testModel{
			ID:  "id-1",
			Int: 1,
			Str: "add",
		})

		deleteEntry := testModel{
			ID:  "id-2",
			Int: 2,
			Str: "remove",
		}
		uow.RegisterInsert(deleteEntry)
		uow.RegisterDelete(deleteEntry)

		uow.RegisterInsert(testModel{
			ID:  "id-3",
			Int: 3,
			Str: "update",
		})
		uow.RegisterUpdate(testModel{
			ID:  "id-3",
			Int: 4,
			Str: "update-1",
		})

		err = uow.Commit()
		assert.NoError(t, err)

		cur, err := c.Find(testDriverFactory.Ctx, bson.M{}, options.Find())
		assert.NoError(t, err)

		rows := make([]testModel, 0)
		for cur.Next(testDriverFactory.Ctx) {
			row := testModel{}
			cur.Decode(&row)
			rows = append(rows, row)
		}
		assert.EqualValues(t, rows, []testModel{
			{
				ID:  "id-1",
				Int: 1,
				Str: "add",
			}, {
				ID:  "id-3",
				Int: 4,
				Str: "update-1",
			},
		})
	})
}
