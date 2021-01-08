package mysqldb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
)

var modelTypeOfSQLMaker = make(map[reflect.Type]*sqlMaker)

type sqlMaker struct {
	Table identity.IStruct

	addSQL    string
	removeSQL string
	saveSQL   string
}

func (m *sqlMaker) GetAdd() (string, error) {
	if m.addSQL != "" {
		return m.addSQL, nil
	}

	tableName, err := m.Table.GetName()
	if err != nil {
		return "", err
	}

	bf := make([]string, 0)
	bf = append(bf, "INSERT INTO `", tableName, "`(")
	columns := m.Table.FindFields()
	underscore.Chain(columns).Each(func(r identity.IField, i int) {
		if i > 0 {
			bf = append(bf, ", ")
		}

		bf = append(
			bf,
			"`",
			r.GetName(),
			"`",
		)
	})
	bf = append(bf, ") VALUES (")
	underscore.Chain(columns).Each(func(r identity.IField, i int) {
		if i > 0 {
			bf = append(bf, ", ")
		}

		bf = append(bf, "?")
	})
	bf = append(bf, ")")
	m.addSQL = strings.Join(bf, "")
	return m.addSQL, nil
}

func (m sqlMaker) GetCount(opt queryOption) (string, error) {
	tableName, err := m.Table.GetName()
	if err != nil {
		return "", err
	}

	bf := make([]string, 0)
	bf = append(bf, "SELECT COUNT(1) FROM `", tableName, "`")
	if opt.Where != "" {
		bf = append(
			bf,
			" WHERE ",
			opt.Where,
		)
	}
	return strings.Join(bf, ""), nil
}

func (m *sqlMaker) GetRemove() (string, error) {
	if m.removeSQL != "" {
		return m.removeSQL, nil
	}

	pk, err := m.Table.GetIDField()
	if err != nil {
		return "", err
	}

	m.removeSQL = fmt.Sprintf(
		"DELETE FROM `%s` WHERE `%s` = ?",
		pk.GetStructName(),
		pk.GetName(),
	)
	return m.removeSQL, nil
}

func (m *sqlMaker) GetSave() (string, error) {
	if m.saveSQL != "" {
		return m.saveSQL, nil
	}

	pk, err := m.Table.GetIDField()
	if err != nil {
		return "", err
	}

	bf := make([]string, 0)
	bf = append(
		bf,
		"UPDATE `",
		pk.GetStructName(),
		"` SET ",
	)
	columns := m.Table.FindFields()
	isFirst := true
	underscore.Chain(columns).Each(func(r identity.IField, _ int) {
		if r.GetStructName() != "" {
			return
		}

		if isFirst {
			isFirst = false
		} else {
			bf = append(bf, ", ")
		}

		bf = append(
			bf,
			"`",
			r.GetName(),
			"` = ?",
		)
	})
	bf = append(
		bf,
		" WHERE `",
		pk.GetName(),
		"` = ?",
	)
	m.saveSQL = strings.Join(bf, "")
	return m.saveSQL, nil
}

func (m sqlMaker) GetSelect(opt queryOption) (string, error) {
	tableName, err := m.Table.GetName()
	if err != nil {
		return "", err
	}

	bf := make([]string, 0)
	bf = append(bf, "SELECT * FROM `", tableName, "`")
	if opt.Where != "" {
		bf = append(bf, " WHERE ", opt.Where)
	}
	if opt.Skip >= 0 && opt.Take > 0 {
		bf = append(
			bf,
			" LIMIT ",
			strconv.Itoa(opt.Skip),
			",",
			strconv.Itoa(opt.Take),
		)
	}
	if len(opt.Orders) > 0 {
		bf = append(bf, " ORDER BY ")
		underscore.Chain(opt.Orders).Each(func(r orderOption, i int) {
			if i > 0 {
				bf = append(bf, ", ")
			}

			bf = append(bf, "`", r.Field, "` ", r.Flag)
		})
	}
	return strings.Join(bf, ""), nil
}

func newSQLMaker(modelType reflect.Type) *sqlMaker {
	if _, ok := modelTypeOfSQLMaker[modelType]; !ok {
		modelTypeOfSQLMaker[modelType] = &sqlMaker{
			Table: identity.NewStruct(modelType),
		}
	}

	return modelTypeOfSQLMaker[modelType]
}
