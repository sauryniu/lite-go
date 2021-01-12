package identity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	ID   string `db:"id,test"`
	Name string `db:"name"`
}

func Test_identityStruct_FindFields(t *testing.T) {
	structType := reflect.TypeOf(testStruct{})
	fields := NewStruct(structType).FindFields()
	defer delete(structTypeOfStruct, structType)

	assert.Len(t, fields, 2)
}

func Test_identityStruct_GetIDField(t *testing.T) {
	structType := reflect.TypeOf(testStruct{})
	idField, err := NewStruct(structType).GetIDField()
	defer delete(structTypeOfStruct, structType)

	assert.NoError(t, err)
	assert.Equal(
		t,
		idField.GetName(),
		"id",
	)
	assert.Equal(
		t,
		idField.GetStructName(),
		"test",
	)
}

type testMissIDStruct struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func Test_identityStruct_GetIDField_Error(t *testing.T) {
	structType := reflect.TypeOf(testMissIDStruct{})
	_, err := NewStruct(structType).GetIDField()
	defer delete(structTypeOfStruct, structType)

	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(
			`缺少^db:"主键,表名"^: %s`,
			structType.Name(),
		),
	)
}

func Test_identityStruct_GetName(t *testing.T) {
	structType := reflect.TypeOf(testStruct{})
	res, err := NewStruct(structType).GetName()
	defer delete(structTypeOfStruct, structType)

	assert.NoError(t, err)
	assert.Equal(t, res, "test")
}

func Test_identityStruct_GetName_Error(t *testing.T) {
	structType := reflect.TypeOf(testMissIDStruct{})
	res, err := NewStruct(structType).GetName()
	defer delete(structTypeOfStruct, structType)

	assert.Empty(t, res)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(
			`缺少^db:"主键,表名"^: %s`,
			structType.Name(),
		),
	)
}

func Test_identityStruct_GetType(t *testing.T) {
	structType := reflect.TypeOf(testMissIDStruct{})
	res := NewStruct(structType).GetType()
	assert.Equal(t, res, structType)
}
