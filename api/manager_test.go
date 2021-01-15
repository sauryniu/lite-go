package api

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_Invalid(t *testing.T) {
	endpoint := "endpoint"
	name := "name"
	res := New(endpoint, name)
	assert.Equal(t, res, invalid)
}

func Test_Register(t *testing.T) {
	endpoint := "endpoint"
	name := "name"
	Register(endpoint, name, invalid)
	defer delete(metadatas, endpoint)

	apiTypes, ok := metadatas[endpoint]
	assert.True(t, ok)

	apiType, ok := apiTypes[name]
	assert.True(t, ok)
	assert.Equal(
		t,
		apiType,
		reflect.TypeOf(invalidAPI{}),
	)
}
