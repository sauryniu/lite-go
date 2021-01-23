package yamlex

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/ioex"
	"github.com/ahl5esoft/lite-go/osex"
	"gopkg.in/yaml.v2"
)

type env map[interface{}]interface{}

func (m env) Get(k string, v interface{}) {
	if value, ok := m[k]; ok {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(value),
		)
	}
}

// NewEnv is yaml IEnv
func NewEnv(bf []byte) osex.IEnv {
	var hash map[interface{}]interface{}
	yaml.Unmarshal(bf, &hash)
	return env(hash)
}

// NewEnvFromFile is yaml IEnv
func NewEnvFromFile(file ioex.IFile) osex.IEnv {
	var bf []byte
	file.Read(&bf)
	return NewEnv(bf)
}
