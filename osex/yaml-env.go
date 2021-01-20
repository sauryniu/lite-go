package osex

import (
	reflect "reflect"

	"github.com/ahl5esoft/lite-go/ioex"
)

type yamlEnv map[interface{}]interface{}

func (m yamlEnv) Get(k string, v interface{}) {
	if value, ok := m[k]; ok {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(value),
		)
	}
}

// NewYamlEnv is yaml IEnv
func NewYamlEnv(yaml map[interface{}]interface{}) IEnv {
	env := yamlEnv(yaml)
	return env
}

// NewYamlEnvFromFile is json IEnv
func NewYamlEnvFromFile(file ioex.IFile) IEnv {
	env := new(yamlEnv)
	file.ReadYaml(env)
	return *env
}
