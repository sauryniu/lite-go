package osex

import (
	"testing"
	"time"

	"github.com/go-playground/assert"
)

func Test_osNowTime_Unix(t *testing.T) {
	assert.Equal(
		t,
		new(NowTime).Unix(),
		time.Now().Unix(),
	)
}

func Test_osNowTime_UnixNano(t *testing.T) {
	assert.Equal(
		t,
		new(NowTime).UnixNano(),
		time.Now().UnixNano(),
	)
}
