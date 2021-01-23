package redisex

import (
	"errors"
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
		gomock.Eq("lock-ok"),
		gomock.Eq(""),
		gomock.Eq("ex"),
		gomock.Eq(0*time.Second),
		gomock.Eq("nx"),
	).Return(true, nil)

	res, resErr := NewLock(redis).Lock("lock-%s", "ok")
	assert.NoError(t, resErr)
	assert.NotNil(t, res)

	redis.EXPECT().Del(
		gomock.Eq("lock-ok"),
	).Return(0, nil)
	res()
}

func Test_redisLock_Lock_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("err")
	redis := NewMockIRedis(ctrl)
	redis.EXPECT().Set(
		gomock.Eq("lock-err"),
		gomock.Eq(""),
		gomock.Eq("ex"),
		gomock.Eq(0*time.Second),
		gomock.Eq("nx"),
	).Return(
		true,
		err,
	)

	res, resErr := NewLock(redis).Lock("lock-err")
	assert.Error(t, resErr)
	assert.Equal(t, resErr, err)
	assert.Nil(t, res)
}

func Test_redisLock_Lock_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redis := NewMockIRedis(ctrl)
	redis.EXPECT().Set(
		gomock.Eq("lock-fail"),
		gomock.Eq(""),
		gomock.Eq("ex"),
		gomock.Eq(0*time.Second),
		gomock.Eq("nx"),
	).Return(false, nil)

	res, resErr := NewLock(redis).Lock("lock-fail")
	assert.NoError(t, resErr)
	assert.Nil(t, res)
}

func Test_redisLock_SetExpire(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redis := NewMockIRedis(ctrl)
	redis.EXPECT().Set(
		gomock.Eq("lock-expires"),
		gomock.Eq(""),
		gomock.Eq("ex"),
		gomock.Eq(5*time.Second),
		gomock.Eq("nx"),
	).Return(true, nil)

	res, resErr := NewLock(redis).SetExpire(5 * time.Second).Lock("lock-expires")
	assert.NoError(t, resErr)
	assert.NotNil(t, res)
}
