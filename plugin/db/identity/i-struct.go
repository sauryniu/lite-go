package identity

import "reflect"

// IStruct is 结构接口
type IStruct interface {
	FindFields() []IField
	GetIDField() (IField, error)
	GetName() (string, error)
	GetType() reflect.Type
}
