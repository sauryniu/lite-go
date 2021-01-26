package yamlex

import (
	"fmt"
	"testing"

	"github.com/ahl5esoft/lite-go/ioex"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NewEnv_Get_string(t *testing.T) {
	key := "str"
	value := "string"
	var res string
	NewEnv(
		[]byte(`str: string`),
	).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_NewEnv_Get_Array(t *testing.T) {
	key := "arr"
	var res []interface{}
	NewEnv(
		[]byte(`arr:
- 1
- 2
- 3`),
	).Get(key, &res)
	assert.EqualValues(t, res, []interface{}{1, 2, 3})
}

func Test_NewEnvFromFile_Get_string(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := "str"
	value := "string"
	text := fmt.Sprintf("%s: %s", key, value)
	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(
		0,
		[]byte(text),
	).Return(nil)

	var res string
	NewEnvFromFile(mockFile).Get(key, &res)
	assert.Equal(t, res, value)
}

func Test_NewEnvFromFile_Get_Array(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := "arr"
	mockFile := ioex.NewMockIFile(ctrl)
	mockFile.EXPECT().Read(
		gomock.Not(nil),
	).SetArg(
		0,
		[]byte(`arr:
- 1
- 2
- 3`),
	).Return(nil)

	var res []interface{}
	NewEnvFromFile(mockFile).Get(key, &res)
	assert.EqualValues(t, res, []interface{}{1, 2, 3})
}
