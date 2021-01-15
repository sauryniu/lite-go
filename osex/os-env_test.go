package osex

import (
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func Test_osEnv_Get_string(t *testing.T) {
	k := "GOPATH"
	var res string
	NewOSEnv().Get(k, &res)
	assert.Equal(
		t,
		res,
		os.Getenv(k),
	)
}

func Test_osEnv_Get_Array(t *testing.T) {
	k := "test-get-json"
	src := []int{1, 2, 3}
	s, _ := jsoniter.MarshalToString(src)
	os.Setenv(k, s)

	var res []int
	NewOSEnv().Get(k, &res)
	assert.EqualValues(t, res, src)
}
