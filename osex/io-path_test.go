package osex

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_osPath_Join(t *testing.T) {
	res := new(osPath).Join("a", "b")
	assert.Equal(
		t,
		res,
		filepath.Join("a", "b"),
	)
}

func Test_osPath_Join_HasParent(t *testing.T) {
	res := new(osPath).Join("a", "b", "..", "c")
	assert.Equal(
		t,
		res,
		filepath.Join("a", "c"),
	)
}
