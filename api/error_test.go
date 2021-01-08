package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewError(t *testing.T) {
	err := NewError(5, "")
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

		err, ok := rc.(CustomError)
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
