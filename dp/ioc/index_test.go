package ioc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ahl5esoft/lite-go/reflectex"
	"github.com/stretchr/testify/assert"
)

type iInterface interface {
	Test()
}

type derive struct{}

func (m derive) Test() {
	fmt.Println("set test")
}

type myTest struct {
	One iInterface `inject:""`
}

func Test_Get(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := getInterfaceType(
		(*iInterface)(nil),
	)
	typeOfInstance[ct] = 1
	defer delete(typeOfInstance, ct)

	res := Get(ct)
	assert.Equal(t, res, 1)
}

func Test_Get_无效类型(t *testing.T) {
	ct := getInterfaceType(
		(*iInterface)(nil),
	)
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(invalidTypeFormat, ct),
		)
	}()

	Get(ct)
}

func Test_Has(t *testing.T) {
	ct := getInterfaceType(
		(*iInterface)(nil),
	)
	typeOfInstance[ct] = 1
	defer delete(typeOfInstance, ct)

	assert.True(
		t,
		Has(ct),
	)
}

func Test_Inject(t *testing.T) {
	it := getInterfaceType(
		(*iInterface)(nil),
	)
	typeOfInstance[it] = new(derive)

	var m myTest
	Inject(&m)

	assert.Equal(t, m.One, typeOfInstance[it])
}

func Test_Remove(t *testing.T) {
	it := getInterfaceType(
		(*iInterface)(nil),
	)
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	Remove(it)
}

func Test_Set(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := reflectex.InterfaceTypeOf(
		(*iInterface)(nil),
	)
	defer delete(typeOfInstance, ct)

	Set(
		ct,
		new(derive),
	)
}

func Test_Set_非接口类型(t *testing.T) {
	it := reflect.TypeOf(1)
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(notInterfaceTypeFormat, it),
		)
	}()
	Set(
		it,
		new(derive),
	)
}

func Test_Set_没有继承(t *testing.T) {
	it := getInterfaceType(
		(*iInterface)(nil),
	)
	v := ""
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(notImplementsFormat, v, it),
		)
	}()
	Set(it, v)
}

func Test_Set_对象(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := reflectex.InterfaceTypeOf((*iInterface)(nil))
	defer delete(typeOfInstance, ct)

	Set(
		(*iInterface)(nil),
		new(derive),
	)

	_, ok := typeOfInstance[ct]
	assert.True(t, ok)
}
