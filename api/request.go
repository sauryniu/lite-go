package api

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/dp/ioc"
)

var (
	invalid   = new(invalidAPI)
	metadatas = make(map[string]map[string]reflect.Type)
)

// ICreateContext is 请求处理上下文
type ICreateContext interface {
	GetAPIName() string
	GetEndpoint() string
	SetAPI(IAPI)
}

// NewCreateHandler is 创建api处理器
func NewCreateHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if rCtx, ok := ctx.(ICreateContext); ok {
			var instance IAPI
			if apiTypes, ok := metadatas[rCtx.GetEndpoint()]; ok {
				if apiType, ok := apiTypes[rCtx.GetAPIName()]; ok {
					instance = reflect.New(apiType).Interface().(IAPI)
					ioc.Inject(instance)
					rCtx.SetAPI(instance)
					return nil
				}
			}
		}

		return NewError(APIErrorCode, "")
	})
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
