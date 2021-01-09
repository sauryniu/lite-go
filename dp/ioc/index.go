package ioc

import (
	"fmt"
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
)

const (
	instanceIsNotPtr = "ioc: 注入实例必须是指针"
	invalidIDFormat  = "ioc: 无效标识(%s)"
)

var idOfInstance = make(map[string]interface{})

// Get is 获取实例
func Get(id string) interface{} {
	if v, ok := idOfInstance[id]; ok {
		return v
	}

	panic(
		fmt.Sprintf(invalidIDFormat, id),
	)
}

// Has is 是否存在
func Has(id string) bool {
	_, ok := idOfInstance[id]
	return ok
}

// Inject is 遍历实例内的需要依赖注入的字段进行注入
func Inject(instance interface{}) {
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() != reflect.Ptr {
		panic(instanceIsNotPtr)
	}

	instanceValue = instanceValue.Elem()
	underscore.Range(
		0,
		instanceValue.Type().NumField(),
		1,
	).Each(func(r int, _ int) {
		field := instanceValue.Type().Field(r)
		id, ok := field.Tag.Lookup("inject")
		if !ok {
			return
		}

		fieldValue := instanceValue.FieldByIndex(field.Index)
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue.Set(
				reflect.New(
					field.Type.Elem(),
				),
			)
			fieldValue = fieldValue.Elem()
		}

		v := Get(id)
		fieldValue.Set(
			reflect.ValueOf(v),
		)
	})
}

// Remove is 删除
func Remove(id string) {
	if Has(id) {
		delete(idOfInstance, id)
	}
}

// Set is 设置依赖注入
func Set(id string, instance interface{}) {
	idOfInstance[id] = instance
}
