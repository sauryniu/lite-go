package api

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testAPI struct{}

func (m testAPI) Auth() bool {
	return false
}

func (m testAPI) Call() interface{} {
	return nil
}

func (m testAPI) Valid(ctx interface{}) bool {
	return false
}

func Test_New(t *testing.T) {
	res := New("a", "b")
	assert.Equal(t, res, invalid)
}

func Test_Register(t *testing.T) {
	Register("a", "aa", testAPI{})

	res := New("a", "aa")
	assert.Equal(
		t,
		reflect.TypeOf(res).Elem(),
		reflect.TypeOf(testAPI{}),
	)
}
