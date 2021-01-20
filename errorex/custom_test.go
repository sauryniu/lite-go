package errorex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	err := New(5, "")
	assert.Equal(
		t,
		int(err.Code),
		5,
	)
	assert.Equal(
		t,
		err.Error(),
		"",
	)
}

func Test_Throw(t *testing.T) {
	defer func() {
		rc := recover()
		assert.NotNil(t, rc)

		err, ok := rc.(Custom)
		assert.True(t, ok)
		assert.Equal(
			t,
			int(err.Code),
			55,
		)
		assert.Equal(
			t,
			err.Error(),
			"panic",
		)
	}()

	Throw(55, "panic")
}
