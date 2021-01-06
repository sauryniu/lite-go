package iopath

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Join(t *testing.T) {
	pathArgs := []string{"a", "b"}
	res := Join(pathArgs...)
	assert.Equal(
		t,
		res,
		filepath.Join(pathArgs...),
	)
}

func Test_Join_HasParent(t *testing.T) {
	res := Join("a", "b", "..", "c")
	assert.Equal(
		t,
		res,
		filepath.Join("a", "c"),
	)
}
