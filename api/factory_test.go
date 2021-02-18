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
		res := factoryInstance.Build(endpoint, name)
		assert.Equal(t, res, invalid)
	})
}

func Test_Register(t *testing.T) {
	endpoint := "endpoint"
	name := "name"
	Register(endpoint, name, invalid)
	defer delete(factoryInstance, endpoint)

	apiTypes, ok := factoryInstance[endpoint]
	assert.True(t, ok)

	apiType, ok := apiTypes[name]
	assert.True(t, ok)
	assert.Equal(
		t,
		apiType,
		reflect.TypeOf(invalidAPI{}),
	)
}
