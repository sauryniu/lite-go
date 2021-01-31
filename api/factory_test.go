package api

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factory_Build(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		endpoint := "endpoint"
		name := "name"
		res := factoryInstane.Build(endpoint, name)
		assert.Equal(t, res, invalid)
	})
}

func Test_Register(t *testing.T) {
	endpoint := "endpoint"
	name := "name"
	Register(endpoint, name, invalid)
	defer delete(factoryInstane, endpoint)

	apiTypes, ok := factoryInstane[endpoint]
	assert.True(t, ok)

	apiType, ok := apiTypes[name]
	assert.True(t, ok)
	assert.Equal(
		t,
		apiType,
		reflect.TypeOf(invalidAPI{}),
	)
}
