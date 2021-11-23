package mongosvc

import (
	"reflect"
)

type fieldMetadata struct {
	columnName string
	field      reflect.StructField
	modelType  reflect.Type
	tableName  *string
}

func (m *fieldMetadata) GetColumnName() string {
	if m.columnName == "" {
		m.columnName = m.field.Tag.Get("db")
		if m.columnName == "" {
			m.columnName = m.field.Name
		}
	}

	return m.columnName
}

func (m *fieldMetadata) GetTableName() string {
	if m.tableName == nil {
		v, ok := m.field.Tag.Lookup("alias")
		if ok && v == "" {
			v = m.modelType.Name()
		}

		m.tableName = &v
	}

	return *m.tableName
}

func (m *fieldMetadata) GetValue(tableValue reflect.Value) interface{} {
	return tableValue.FieldByIndex(m.field.Index).Interface()
}

func newFieldMetadata(field reflect.StructField, modelType reflect.Type) *fieldMetadata {
	return &fieldMetadata{
		field:     field,
		modelType: modelType,
	}
}
