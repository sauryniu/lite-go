package osex

import (
	"testing"
	"time"

	"github.com/go-playground/assert"
)

func Test_osNowTime_Unix(t *testing.T) {
	assert.Equal(
		t,
		NewNowUnix().Unix(),
		time.Now().Unix(),
	)
}

func Test_osNowTime_UnixNano(t *testing.T) {
	assert.Equal(
		t,
		NewNowUnixNano().UnixNano(),
		time.Now().UnixNano(),
	)
}
