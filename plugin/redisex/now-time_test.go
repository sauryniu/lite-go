package redisex

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_NowTime_Unix(t *testing.T) {
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

func Test_NowTime_UnixNano(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := NewMockIRedis(ctrl)
	mockRedis.EXPECT().Time().Return(
		time.Unix(100, 0),
		nil,
	)

	res := NewNowTime(mockRedis).UnixNano()
	assert.NotEqual(t, res, 100)
}
