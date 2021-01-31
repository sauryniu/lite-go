package api

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/dp/ioc"
)

var (
	factoryInstane = make(factory)
	invalid        = new(invalidAPI)
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
	return factoryInstane
}

// Register is 注册api
func Register(endpoint, name string, api IAPI) {
	if _, ok := factoryInstane[endpoint]; !ok {
		factoryInstane[endpoint] = make(map[string]reflect.Type)
	}

	apiType := reflect.TypeOf(api)
	if apiType.Kind() == reflect.Ptr {
		apiType = apiType.Elem()
	}
	factoryInstane[endpoint][name] = apiType
}
