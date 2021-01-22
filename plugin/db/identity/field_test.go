package identity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testField struct {
	FieldA string `db:"a" alias:""`
	FieldB string `db:"b" alias:"test"`
	FieldC string
}

func Test_field_GetName(t *testing.T) {
	structType := reflect.TypeOf(testField{})
	f, _ := structType.FieldByName("FieldA")
	self := NewField(f, structType)
	assert.Equal(
		t,
		self.GetName(),
		"a",
	)
}

func Test_field_GetName_NoTag(t *testing.T) {
	structType := reflect.TypeOf(testField{})
	f, _ := structType.FieldByName("FieldC")
	self := NewField(f, structType)
	assert.Equal(
		t,
		self.GetName(),
		"FieldC",
	)
}

func Test_field_GetStructName_元数据(t *testing.T) {
	structType := reflect.TypeOf(testField{})
	f, _ := structType.FieldByName("FieldB")
	self := NewField(f, structType)
	assert.Equal(
		t,
		self.GetStructName(),
		"test",
	)
}

func Test_field_GetStructName_模型名(t *testing.T) {
	structType := reflect.TypeOf(testField{})
	f, _ := structType.FieldByName("FieldA")
	self := NewField(f, structType)
	assert.Equal(
		t,
		self.GetStructName(),
		structType.Name(),
	)
}

func Test_field_GetValue(t *testing.T) {
	tf := testField{
		FieldA: "aa",
	}
	tfValue := reflect.ValueOf(tf)
	f, _ := tfValue.Type().FieldByName("FieldA")
	self := NewField(
		f,
		tfValue.Type(),
	)
	assert.Equal(
		t,
		self.GetValue(tfValue),
		tf.FieldA,
	)
}
