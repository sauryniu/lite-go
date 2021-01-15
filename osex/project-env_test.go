package osex

import (
	"os"
	"testing"

	"github.com/ahl5esoft/lite-go/ioex"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_projectEnv_Get_string(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIFactory := ioex.NewMockIFactory(ctrl)

	wd, _ := os.Getwd()
	mockDir := ioex.NewMockIDirectory(ctrl)
	mockIFactory.EXPECT().BuildDirectory(
		gomock.Eq(wd),
	).Return(mockDir)

	mockDir.EXPECT().GetName().Return("test-go")

	mockEnv := NewMockIEnv(ctrl)
	mockEnv.EXPECT().Get(
		gomock.Eq("test_go_string"),
		gomock.Not(nil),
	).SetArg(1, "ok").Return()

	var res string
	NewProjectEnv(mockIFactory, mockEnv).Get("string", &res)
	assert.Equal(t, res, "ok")
}

func Test_projectEnv_Get_array(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIFactory := ioex.NewMockIFactory(ctrl)

	wd, _ := os.Getwd()
	mockDir := ioex.NewMockIDirectory(ctrl)
	mockIFactory.EXPECT().BuildDirectory(
		gomock.Eq(wd),
	).Return(mockDir)

	mockDir.EXPECT().GetName().Return("test-go")

	mockEnv := NewMockIEnv(ctrl)
	mockEnv.EXPECT().Get(
		gomock.Eq("test_go_arr"),
		gomock.Not(nil),
	).SetArg(1, []int{1, 2}).Return()

	var res []int
	NewProjectEnv(mockIFactory, mockEnv).Get("arr", &res)
	assert.EqualValues(t, res, []int{1, 2})
}
