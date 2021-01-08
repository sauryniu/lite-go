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
	fields := NewStruct(
		reflect.TypeOf(testStruct{}),
	).FindFields()
	assert.Len(t, fields, 2)
}

func Test_identityStruct_GetIDField(t *testing.T) {
	idField, err := NewStruct(
		reflect.TypeOf(testStruct{}),
	).GetIDField()
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
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(
			"缺少标识: %s",
			structType.Name(),
		),
	)
}

func Test_identityStruct_GetName(t *testing.T) {
	res, err := NewStruct(
		reflect.TypeOf(testStruct{}),
	).GetName()
	assert.NoError(t, err)
	assert.Equal(t, res, "test")
}

func Test_identityStruct_GetName_Error(t *testing.T) {
	structType := reflect.TypeOf(testMissIDStruct{})
	res, err := NewStruct(structType).GetName()
	assert.Empty(t, res)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(
			"缺少标识: %s",
			structType.Name(),
		),
	)
}

func Test_identityStruct_GetType(t *testing.T) {
	structType := reflect.TypeOf(testMissIDStruct{})
	res := NewStruct(structType).GetType()
	assert.Equal(t, res, structType)
}
