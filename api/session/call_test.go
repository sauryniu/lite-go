package session

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_caller_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(sessionCaller)
	self.getRoute = "get"
	key := "key"

	mockCaller := api.NewMockICaller(ctrl)
	self.ICaller = mockCaller

	mockCaller.EXPECT().Call(self.getRoute, getMessage{
		Key: key,
	}).Return(`{"name":"hello"}`, nil)

	var s struct {
		Name string
	}
	err := self.Get(key, &s)
	assert.NoError(t, err)
	assert.Equal(t, s.Name, "hello")
}

func Test_caller_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(sessionCaller)
	self.setRoute = "sr"
	body := []int{1, 2, 3}
	expires := 5 * time.Second
	key := "key"

	mockCaller := api.NewMockICaller(ctrl)
	self.ICaller = mockCaller

	mockCaller.EXPECT().Call(self.setRoute, setMessage{
		Expires:  5,
		Interval: 0,
		Value:    `[1,2,3]`,
	}).Return(key, nil)

	res, err := self.Set(body, expires, time.Nanosecond)
	assert.NoError(t, err)
	assert.Equal(t, res, key)
}
