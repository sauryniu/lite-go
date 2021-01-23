package redisex

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/timeex"
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

	var nowUnix timeex.INowUnix
	nowUnix = &NowTime{mockRedis}
	res := nowUnix.Unix()
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

	var nowUnixNano timeex.INowUnixNano
	nowUnixNano = &NowTime{mockRedis}
	res := nowUnixNano.UnixNano()
	assert.NotEqual(t, res, 100)
}
