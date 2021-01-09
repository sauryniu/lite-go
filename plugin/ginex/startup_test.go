package ginex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewStartup(t *testing.T) {
	err := NewStartup().Handle(nil)
	assert.NoError(t, err)
}
