package identity

import (
	"fmt"
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
)

var structTypeOfStruct = make(map[reflect.Type]IStruct)

type identityStruct struct {
	fields     []IField
	idIndex    int
	structType reflect.Type
}

func (m *identityStruct) FindFields() []IField {
	if m.fields == nil {
		m.idIndex = -1
		underscore.Range(
			0,
			m.structType.NumField(),
			1,
		).Map(func(r int, i int) IField {
			c := NewField(
				m.structType.Field(r),
				m.structType,
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
			m.structType.Name(),
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
	return m.structType
}

// NewStruct is IStruct实例
func NewStruct(structType reflect.Type) IStruct {
	if _, ok := structTypeOfStruct[structType]; !ok {
		structTypeOfStruct[structType] = &identityStruct{
			structType: structType,
		}
	}

	return structTypeOfStruct[structType]
}
