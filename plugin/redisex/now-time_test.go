package redisex

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_redisNowTime_NanoUnix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := NewMockIRedis(ctrl)
	mockRedis.EXPECT().Time().Return(
		time.Unix(100, 0),
		nil,
	)

	res := NewNowTime(mockRedis).NanoUnix()
	assert.NotEqual(t, res, 100)
}

func Test_redisNowTime_Unix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := NewMockIRedis(ctrl)
	mockRedis.EXPECT().Time().Return(
		time.Unix(100, 0),
		nil,
	)

	res := NewNowTime(mockRedis).Unix()
	assert.NotEqual(t, res, 100)
}
