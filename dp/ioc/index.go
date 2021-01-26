package ioc

import (
	"fmt"
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
)

const (
	instanceIsNotPtr       = "ioc: 注入实例必须是指针"
	invalidTypeFormat      = "ioc: 无效类型(%v)"
	notImplementsFormat    = "ioc: %v没有实现%v"
	notInterfaceTypeFormat = "ioc: 非接口类型(%v)"
)

var typeOfInstance = make(map[reflect.Type]interface{})

// Get is 获取实例
func Get(interfaceObj interface{}) interface{} {
	interfaceType := getInterfaceType(interfaceObj)
	if v, ok := typeOfInstance[interfaceType]; ok {
		return v
	}

	panic(
		fmt.Errorf(invalidTypeFormat, interfaceType),
	)
}

// Has is 是否存在
func Has(interfaceObj interface{}) bool {
	interfaceType := getInterfaceType(interfaceObj)
	_, ok := typeOfInstance[interfaceType]
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
		_, ok := field.Tag.Lookup("inject")
		if !ok {
			return
		}

		fieldValue := instanceValue.FieldByIndex(field.Index)
		if fieldValue.Kind() == reflect.Ptr {
			value := reflect.New(
				field.Type.Elem(),
			)
			fieldValue.Set(value)
			fieldValue = fieldValue.Elem()
		}

		v := Get(field.Type)
		fieldValue.Set(
			reflect.ValueOf(v),
		)
	})
}

// Remove is 删除
func Remove(t reflect.Type) {
	if Has(t) {
		delete(typeOfInstance, t)
	}
}

// Set is 设置依赖注入
func Set(interfaceObj interface{}, instance interface{}) {
	interfaceType := getInterfaceType(interfaceObj)
	instanceType := reflect.TypeOf(instance)
	if !instanceType.Implements(interfaceType) {
		panic(
			fmt.Errorf(notImplementsFormat, instance, interfaceType),
		)
	}

	typeOfInstance[interfaceType] = instance
}

func getInterfaceType(interfaceObj interface{}) reflect.Type {
	var interfaceType reflect.Type
	var ok bool
	if interfaceType, ok = interfaceObj.(reflect.Type); !ok {
		interfaceType = reflect.TypeOf(interfaceObj)
	}

	if interfaceType.Kind() == reflect.Ptr {
		interfaceType = interfaceType.Elem()
	}

	if interfaceType.Kind() != reflect.Interface {
		panic(
			fmt.Errorf(notInterfaceTypeFormat, interfaceType),
		)
	}

	return interfaceType
}
