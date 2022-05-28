package mongosvc

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testModelMetadataModel struct {
	ID   string `alias:"test" db:"id"`
	Name string `db:"name"`
}

type testMissIDModelMetadataModel struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func Test_modelMetadata_FindFields(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := getModelMetadata(
			reflect.TypeOf(testModelMetadataModel{}),
		)
		defer modelMetadatas.Delete(self.modelType)

		fields := self.FindFields()
		assert.Len(t, fields, 2)
	})
}

func Test_modelMetadata_GetIDField(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := getModelMetadata(
			reflect.TypeOf(testModelMetadataModel{}),
		)
		defer modelMetadatas.Delete(self.modelType)

		idField, err := self.GetIDField()
		assert.NoError(t, err)
		assert.Equal(
			t,
			idField.GetColumnName(),
			"id",
		)
		assert.Equal(
			t,
			idField.GetTableName(),
			"test",
		)
	})

	t.Run("err", func(t *testing.T) {
		self := getModelMetadata(
			reflect.TypeOf(testMissIDModelMetadataModel{}),
		)
		defer modelMetadatas.Delete(self.modelType)

		defer func() {
			rv := recover()
			assert.NotNil(t, rv)
		}()

		self.GetIDField()
	})
}

func Test_modelMetadata_GetTableName(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := getModelMetadata(
			reflect.TypeOf(testModelMetadataModel{}),
		)
		defer modelMetadatas.Delete(self.modelType)

		res, err := self.GetTableName()
		assert.NoError(t, err)
		assert.Equal(t, res, "test")
	})

	t.Run("err", func(t *testing.T) {
		self := getModelMetadata(
			reflect.TypeOf(testMissIDModelMetadataModel{}),
		)
		defer func() {
			modelMetadatas.Delete(self.modelType)

			rv := recover()
			assert.NotNil(t, rv)
		}()

		self.GetTableName()
	})
}

func Test_modelMetadata_GetType(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := getModelMetadata(
			reflect.TypeOf(testModelMetadataModel{}),
		)
		defer modelMetadatas.Delete(self.modelType)

		res := self.GetType()
		assert.Equal(t, res, self.modelType)
	})
}
