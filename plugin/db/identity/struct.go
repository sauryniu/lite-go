package identity

import (
	"fmt"
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
)

var structTypeOfStruct = make(map[reflect.Type]IStruct)

type identityStruct struct {
	StructType reflect.Type

	fields  []IField
	idIndex int
}

func (m *identityStruct) FindFields() []IField {
	if m.fields == nil {
		m.idIndex = -1
		underscore.Range(
			0,
			m.StructType.NumField(),
			1,
		).Map(func(r int, i int) IField {
			c := NewField(
				m.StructType.Field(r),
				m.StructType,
			)
			if c.GetStructName() != "" {
				m.idIndex = i
			}
			return c
		}).Value(&m.fields)
	}

	return m.fields
}

func (m *identityStruct) GetIDField() (IField, error) {
	fields := m.FindFields()
	if m.idIndex == -1 {
		return nil, fmt.Errorf(
			`缺少^db:"主键,表名"^: %s`,
			m.StructType.Name(),
		)
	}

	return fields[m.idIndex], nil
}

func (m *identityStruct) GetName() (string, error) {
	idField, err := m.GetIDField()
	if err != nil {
		return "", err
	}

	return idField.GetStructName(), nil
}

func (m *identityStruct) GetType() reflect.Type {
	return m.StructType
}

// NewStruct is IStruct实例
func NewStruct(structType reflect.Type) IStruct {
	if _, ok := structTypeOfStruct[structType]; !ok {
		structTypeOfStruct[structType] = &identityStruct{
			StructType: structType,
		}
	}

	return structTypeOfStruct[structType]
}
