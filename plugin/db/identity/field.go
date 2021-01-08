package identity

import (
	"reflect"
	"strings"
)

type identityField struct {
	Field      reflect.StructField
	Name       string
	StructName string
}

func (m identityField) GetName() string {
	return m.Name
}

func (m identityField) GetStructName() string {
	return m.StructName
}

func (m identityField) GetValue(structValue reflect.Value) interface{} {
	return structValue.FieldByIndex(m.Field.Index).Interface()
}

// NewField is IField实例
func NewField(f reflect.StructField, identityType reflect.Type) IField {
	tag := f.Tag.Get("db")
	tagArgs := strings.Split(tag, ",")

	name := tagArgs[0]
	if name == "" {
		name = f.Name
	}

	var structName string
	if len(tagArgs) > 1 {
		structName = tagArgs[1]
		if tagArgs[1] == "" {
			structName = identityType.Name()
		}
	}
	return identityField{
		Field:      f,
		Name:       name,
		StructName: structName,
	}
}
