package ginex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewStartupHandler(t *testing.T) {
	err := NewStartupHandler().Handle(nil)
	assert.NoError(t, err)
}
