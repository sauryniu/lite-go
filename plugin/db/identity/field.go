package identity

import (
	"reflect"
	"strings"
)

type identityField struct {
	field      reflect.StructField
	name       string
	structName string
}

func (m identityField) GetName() string {
	return m.name
}

func (m identityField) GetStructName() string {
	return m.structName
}

func (m identityField) GetValue(structValue reflect.Value) interface{} {
	return structValue.FieldByIndex(m.field.Index).Interface()
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
		field:      f,
		name:       name,
		structName: structName,
	}
}
