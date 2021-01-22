package identity

import (
	"reflect"
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
	name := f.Tag.Get("db")
	if name == "" {
		name = f.Name
	}

	structName, ok := f.Tag.Lookup("alias")
	if ok && structName == "" {
		structName = identityType.Name()
	}
	return identityField{
		field:      f,
		name:       name,
		structName: structName,
	}
}
