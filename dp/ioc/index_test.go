package ioc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	idOfInstance = make(map[string]interface{})

	id := "pi"
	instance := 3.1415
	Set(id, instance)

	v := Get(id)
	assert.Equal(t, v, instance)
}

func Test_Get_无效(t *testing.T) {
	idOfInstance = make(map[string]interface{})

	id := "pi-1"
	defer func() {
		rv := recover()
		assert.Equal(
			t,
			rv,
			fmt.Sprintf(invalidIDFormat, id),
		)
	}()

	Get(id)
}

func Test_Has_N(t *testing.T) {
	assert.False(
		t,
		Has("Test_Has_N"),
	)
}

func Test_Has_Y(t *testing.T) {
	idOfInstance["Test_Has_Y"] = ""
	assert.True(
		t,
		Has("Test_Has_Y"),
	)

	delete(idOfInstance, "Test_Has_Y")
}

type structStruct struct {
	Pi float64 `inject:"pi"`
}

func Test_Inject_值类型(t *testing.T) {
	idOfInstance = make(map[string]interface{})

	id := "pi"
	instance := 3.1415
	Set(id, instance)

	var v structStruct
	Inject(&v)

	assert.Equal(t, v.Pi, instance)
}

func Test_Inject_值类型_非指针(t *testing.T) {
	idOfInstance = make(map[string]interface{})

	defer func() {
		rv := recover()
		assert.Equal(t, rv, instanceIsNotPtr)
	}()

	id := "pi"
	instance := 3.1415
	Set(id, instance)

	var v structStruct
	Inject(v)
}

type ptrStruct struct {
	Struct *structStruct `inject:"s"`
}

func Test_Inject_指针(t *testing.T) {
	idOfInstance = make(map[string]interface{})

	instance := structStruct{
		Pi: 3.1415,
	}
	Set("s", instance)

	var p ptrStruct
	Inject(&p)

	assert.Equal(t, *p.Struct, instance)
}
