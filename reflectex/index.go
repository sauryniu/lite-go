package reflectex

import (
	"fmt"
	"reflect"
)

var notInterfaceTypeFormat = "非接口类型: %v"

// InterfaceTypeOf is 获取接口类型
func InterfaceTypeOf(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic(
			fmt.Errorf(notInterfaceTypeFormat, i),
		)
	}

	return t
}
