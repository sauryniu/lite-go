package api

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/dp/ioc"
)

var (
	factoryInstance = make(factory)
	invalid         = new(invalidAPI)
)

type factory map[string]map[string]reflect.Type

func (m factory) Build(endpoint, name string) IAPI {
	if apiTypes, ok := m[endpoint]; ok {
		if apiType, ok := apiTypes[name]; ok {
			instance := reflect.New(apiType).Interface().(IAPI)
			ioc.Inject(instance)
			return instance
		}
	}
	return invalid
}

// NewFactory is 创建API工厂
func NewFactory() IFactory {
	return factoryInstance
}

// Register is 注册api
func Register(endpoint, name string, api IAPI) {
	if _, ok := factoryInstance[endpoint]; !ok {
		factoryInstance[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	factoryInstance[endpoint][name] = apiType
}
