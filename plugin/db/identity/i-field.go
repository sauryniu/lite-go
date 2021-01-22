package identity

import "reflect"

// IField is 字段接口
type IField interface {
	GetField() reflect.StructField
	GetName() string
	GetStructName() string
	GetValue(structValue reflect.Value) interface{}
}
