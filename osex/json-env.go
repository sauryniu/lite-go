package osex

import (
	reflect "reflect"

	"github.com/ahl5esoft/lite-go/ioex"
)

type jsonEnv map[string]interface{}

func (m jsonEnv) Get(k string, v interface{}) {
	if value, ok := m[k]; ok {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(value),
		)
	}
}

// NewJSONEnv is json IEnv
func NewJSONEnv(json map[string]interface{}) IEnv {
	env := jsonEnv(json)
	return env
}

// NewJSONEnvFromFile is json IEnv
func NewJSONEnvFromFile(file ioex.IFile) IEnv {
	env := new(jsonEnv)
	file.ReadJSON(env)
	return *env
}
