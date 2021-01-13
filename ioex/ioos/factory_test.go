package ioos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factory_BuildDirectory(t *testing.T) {
	_, ok := new(factory).BuildDirectory().(*directory)
	assert.True(t, ok)
}

func Test_factory_BuildFile(t *testing.T) {
	_, ok := new(factory).BuildFile().(*file)
	assert.True(t, ok)
}
