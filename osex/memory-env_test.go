package osex

import (
	"testing"

	"github.com/ahl5esoft/lite-go/ioex"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_memoryEnv_Get_string(t *testing.T) {
	key := "str"
	value := "string"
	var res string
	NewMemoryEnv(map[string]interface{}{
		key: value,
	}).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_memoryEnv_Get_Array(t *testing.T) {
	key := "arr"
	value := []int{1, 2, 3}
	var res []int
	NewMemoryEnv(map[string]interface{}{
		key: value,
	}).Get(key, &res)
	assert.EqualValues(t, res, value)
}

func Test_memoryEnv_JSON_Get_String(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(
		0,
		[]byte(`{"k":"v"}`),
	).Return(nil)

	var res string
	NewJSONFileEnv(mockFile).Get("k", &res)
	assert.Equal(t, res, "v")
}

func Test_memoryEnv_JSON_Get_Array(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(
		0,
		[]byte(`{"k":[1]}`),
	).Return(nil)

	var res interface{}
	NewJSONFileEnv(mockFile).Get("k", &res)
	assert.EqualValues(
		t,
		res,
		[]interface{}{float64(1)},
	)
}

func Test_memoryEnv_Yaml_Get_String(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(
		0,
		[]byte(`k: v`),
	).Return(nil)

	var res string
	NewYamlFileEnv(mockFile).Get("k", &res)
	assert.Equal(t, res, "v")
}

func Test_memoryEnv_Yaml_Get_Array(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(
		0,
		[]byte(`k: [1]`),
	).Return(nil)

	var res interface{}
	NewYamlFileEnv(mockFile).Get("k", &res)
	assert.EqualValues(
		t,
		res,
		[]interface{}{int(1)},
	)
}
