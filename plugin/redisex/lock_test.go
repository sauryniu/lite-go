package redisex

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_redisLock_Lock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redis := NewMockIRedis(ctrl)
	redis.EXPECT().Set(
		gomock.Eq("test: 5"),
		gomock.Eq(""),
		gomock.Eq("ex"),
		gomock.Eq(0*time.Second),
		gomock.Eq("nx"),
	).Return(true, nil)

	ok, _ := NewLock(redis).Lock("test: %d", 5)
	assert.True(t, ok)
}

func Test_redisLock_SetExpire(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redis := NewMockIRedis(ctrl)
	redis.EXPECT().Set(
		gomock.Eq("test: 5"),
		gomock.Eq(""),
		gomock.Eq("ex"),
		gomock.Eq(5*time.Second),
		gomock.Eq("nx"),
	).Return(true, nil)

	ok, _ := NewLock(redis).SetExpire(5*time.Second).Lock("test: %d", 5)
	assert.True(t, ok)
}

func Test_rediscLock_Unlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redis := NewMockIRedis(ctrl)
	redis.EXPECT().Del(
		gomock.Eq(""),
	).Return(0, nil)

	NewLock(redis).Unlock()
}
