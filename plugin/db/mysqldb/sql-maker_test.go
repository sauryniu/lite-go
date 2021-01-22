package mysqldb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSQLMaker struct {
	ID   string `db:"id,tname"`
	Name string
	Age  int `db:"a"`
}

func Test_GetAdd(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetAdd()
	assert.Equal(t, res, "INSERT INTO `tname`(`id`, `Name`, `a`) VALUES (?, ?, ?)")
}

func Test_sqlMaker_GetCount(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetCount(queryOption{})
	assert.Equal(t, res, "SELECT COUNT(1) FROM `tname`")
}

func Test_sqlMaker_GetCount_条件(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetCount(queryOption{
		Where: "id > ?",
	})
	assert.Equal(t, res, "SELECT COUNT(1) FROM `tname` WHERE id > ?")
}

func Test_sqlMaker_GetRemove(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetRemove()
	assert.Equal(t, res, "DELETE FROM `tname` WHERE `id` = ?")
}

func Test_sqlMaker_GetSave(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetSave()
	assert.Equal(t, res, "UPDATE `tname` SET `Name` = ?, `a` = ? WHERE `id` = ?")
}

func Test_sqlMaker_GetSelect(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetSelect(queryOption{})
	assert.Equal(t, res, "SELECT `id` as `ID`, `Name` as `Name`, `a` as `Age` FROM `tname`")
}

func Test_sqlMaker_GetSelect_排序(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetSelect(queryOption{
		Orders: []orderOption{
			{
				Field: "id",
				Flag:  "DESC",
			},
			{
				Field: "name",
				Flag:  "ASC",
			},
		},
	})
	assert.Equal(t, res, "SELECT `id` as `ID`, `Name` as `Name`, `a` as `Age` FROM `tname` ORDER BY `id` DESC, `name` ASC")
}

func Test_sqlMaker_GetSelect_条件(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetSelect(queryOption{
		Where: "id > 10",
	})
	assert.Equal(t, res, "SELECT `id` as `ID`, `Name` as `Name`, `a` as `Age` FROM `tname` WHERE id > 10")
}

func Test_sqlMaker_GetSelect_Skip或Take大于0(t *testing.T) {
	res, _ := newSQLMaker(
		reflect.TypeOf(testSQLMaker{}),
	).GetSelect(queryOption{
		Skip: 10,
		Take: 100,
	})
	assert.Equal(t, res, "SELECT `id` as `ID`, `Name` as `Name`, `a` as `Age` FROM `tname` LIMIT 10,100")
}
