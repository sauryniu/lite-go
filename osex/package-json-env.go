package osex

import (
	reflect "reflect"

	"github.com/ahl5esoft/lite-go/ioex"
)

type packageJSONEnv struct {
	Hash map[string]interface{}
}

func (m packageJSONEnv) Get(k string, v interface{}) {
	if value, ok := m.Hash[k]; ok {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(value),
		)
	}
}

// NewPackageJSONEnv is 创建package.json IEnv
func NewPackageJSONEnv(file ioex.IFile) IEnv {
	env := new(packageJSONEnv)
	file.Read(&env.Hash)
	return env
}
