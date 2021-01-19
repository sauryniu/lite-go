package osex

import (
	"encoding/json"
	"reflect"

	"github.com/ahl5esoft/lite-go/ioex"
	"gopkg.in/yaml.v2"
)

type memoryEnv map[string]interface{}

func (m memoryEnv) Get(k string, v interface{}) {
	if value, ok := m[k]; ok {
		reflect.ValueOf(v).Elem().Set(
			reflect.ValueOf(value),
		)
	}
}

// NewMemoryEnv is 创建内存IEnv
func NewMemoryEnv(memory map[string]interface{}) IEnv {
	return memoryEnv(memory)
}

// NewJSONFileEnv is 创建IEnv(json文件)
func NewJSONFileEnv(file ioex.IFile) IEnv {
	var bf []byte
	file.Read(&bf)

	env := new(memoryEnv)
	json.Unmarshal(bf, &env)
	return env
}

// NewYamlFileEnv is 创建IEnv(yaml文件)
func NewYamlFileEnv(file ioex.IFile) IEnv {
	var bf []byte
	file.Read(&bf)

	env := new(memoryEnv)
	yaml.Unmarshal(bf, &env)
	return env
}
