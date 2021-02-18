//go:generate mockgen -destination i-field_mock.go -package identity github.com/ahl5esoft/lite-go/db/identity IField

package identity

import "reflect"

// IField is 字段接口
type IField interface {
	GetName() string
	GetStructName() string
	GetValue(structValue reflect.Value) interface{}
}
