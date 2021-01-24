package reflectex

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type iTest interface{}

func Test_InterfaceTypeOf(t *testing.T) {
	defer func() {
		assert.Nil(
			t,
			recover(),
		)
	}()

	ct := InterfaceTypeOf((*iTest)(nil))
	assert.Equal(
		t,
		ct,
		reflect.TypeOf((*iTest)(nil)).Elem(),
	)
}

func Test_InterfaceTypeOf_error(t *testing.T) {
	v := (*int)(nil)
	defer func() {
		rv := recover()
		assert.NotNil(t, rv)

		err, ok := rv.(error)
		assert.True(t, ok)
		assert.Equal(
			t,
			err,
			fmt.Errorf(notInterfaceTypeFormat, v),
		)
	}()

	InterfaceTypeOf(v)
}
