package osex

import (
	"os"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

type osEnv struct{}

func (m osEnv) Get(k string, v interface{}) {
	s := os.Getenv(k)
	value := reflect.ValueOf(v).Elem()
	if value.Kind() == reflect.String {
		value.SetString(s)
	} else {
		jsoniter.UnmarshalFromString(s, v)
	}
}

// NewEnv is 创建系统IEnv
func NewEnv() IEnv {
	return new(osEnv)
}
