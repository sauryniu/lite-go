package api

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/dp/ioc"
)

var (
	invalid   = new(invalidAPI)
	metadatas = make(map[string]map[string]reflect.Type)
)

// New is 创建API
func New(endpoint, name string) IAPI {
	if apiTypes, ok := metadatas[endpoint]; ok {
		if apiType, ok := apiTypes[name]; ok {
			instance := reflect.New(apiType).Interface().(IAPI)
			ioc.Inject(instance)
			return instance
		}
	}
	return invalid
}

// Register is 注册api
func Register(endpoint, name string, api IAPI) {
	if _, ok := metadatas[endpoint]; !ok {
		metadatas[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	metadatas[endpoint][name] = apiType
}