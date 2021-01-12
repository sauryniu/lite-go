package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testAPI struct {
}

func (m testAPI) Call() (interface{}, error) {
	return "ok", nil
}

func (m testAPI) GetScope() int {
	return 0
}

type testCreateContext struct {
	API      IAPI
	APIName  string
	Endpoint string
}

func (m testCreateContext) GetEndpoint() string {
	return m.Endpoint
}

func (m testCreateContext) GetAPIName() string {
	return m.APIName
}

func (m *testCreateContext) SetAPI(instance IAPI) {
	m.API = instance
}

func Test_NewCreateHandler(t *testing.T) {
	ctx := &testCreateContext{
		APIName:  "name",
		Endpoint: "endpoint",
	}
	testAPIType := reflect.TypeOf(testAPI{})
	metadatas[ctx.Endpoint] = map[string]reflect.Type{
		ctx.APIName: testAPIType,
	}
	defer delete(metadatas, ctx.Endpoint)

	err := NewCreateHandler().Handle(ctx)
	assert.NoError(t, err)

	resAPIType := reflect.TypeOf(ctx.API)
	fmt.Println("resAPIType", resAPIType)
	assert.Equal(
		t,
		resAPIType.Kind(),
		reflect.Ptr,
	)
	assert.Equal(
		t,
		resAPIType.Elem(),
		testAPIType,
	)
}

func Test_NewCreateHandler_invalid(t *testing.T) {
	ctx := &testCreateContext{
		APIName:  "name",
		Endpoint: "endpoint",
	}
	err := NewCreateHandler().Handle(ctx)
	assert.Error(t, err)
	assert.Equal(
		t,
		err.(CustomError).Code,
		APIErrorCode,
	)
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
