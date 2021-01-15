package osex

import (
	"testing"

	"github.com/ahl5esoft/lite-go/ioex"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_packageJSONEnv_Get_string(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := "str"
	value := "string"
	file := ioex.NewMockIFile(ctrl)
	file.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(0, map[string]interface{}{
		key: value,
	}).Return(nil)

	var res string
	NewPackageJSONEnv(file).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_packageJSONEnv_Get_Array(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := "arr"
	value := []int{1, 2, 3}
	file := ioex.NewMockIFile(ctrl)
	file.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(0, map[string]interface{}{
		key: value,
	}).Return(nil)

	var res []int
	NewPackageJSONEnv(file).Get(key, &res)
	assert.EqualValues(t, res, value)
}
