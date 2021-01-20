package osex

import (
	"testing"

	"github.com/ahl5esoft/lite-go/ioex"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewYamlEnv_Get_string(t *testing.T) {
	key := "str"
	value := "string"
	var res string
	NewYamlEnv(map[interface{}]interface{}{
		key: value,
	}).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_NewYamlEnv_Get_Array(t *testing.T) {
	key := "str"
	value := []int{1, 2, 3}
	var res []int
	NewYamlEnv(map[interface{}]interface{}{
		key: value,
	}).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_NewYamlEnvFromFile_Get_string(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := "str"
	value := "string"
	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().ReadYaml(
		gomock.Not(nil),
	).SetArg(0, map[interface{}]interface{}{
		key: value,
	}).Return(nil)

	var res string
	NewYamlEnvFromFile(mockFile).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_NewYamlEnvFromFile_Get_Array(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := "arr"
	value := []int{1, 2, 3}
	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().ReadYaml(
		gomock.Not(nil),
	).SetArg(0, map[interface{}]interface{}{
		key: value,
	}).Return(nil)

	var res []int
	NewYamlEnvFromFile(mockFile).Get(key, &res)
	assert.EqualValues(t, res, value)
}
