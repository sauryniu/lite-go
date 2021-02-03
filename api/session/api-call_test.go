package session

import (
	"fmt"
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_apiCaller_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(apiCaller)
	self.app = "ss"
	key := "key"

	mockCaller := api.NewMockICaller(ctrl)
	self.ICaller = mockCaller

	mockCaller.EXPECT().Call(
		fmt.Sprintf(getRouteFormat, self.app),
		getMessage{
			Key: key,
		},
		5*time.Second,
	).Return(`{"name":"hello"}`, nil)

	var s struct {
		Name string
	}
	err := self.Get(key, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.Name, "hello")
}

func Test_apiCaller_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(apiCaller)
	self.app = "ss"
	body := []int{1, 2, 3}
	expires := 5 * time.Second
	key := "key"

	mockCaller := api.NewMockICaller(ctrl)
	self.ICaller = mockCaller

	mockCaller.EXPECT().Call(
		fmt.Sprintf(setRouteFormat, self.app),
		setMessage{
			Expires:  5,
			Interval: 0,
			Value:    `[1,2,3]`,
		},
		5*time.Second,
	).Return(key, nil)

	res, err := self.Set(body, expires, time.Nanosecond)
	assert.NoError(t, err)
	assert.Equal(t, res, key)
}
