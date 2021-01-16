package mysqldb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const connString = "root:123456@tcp(10.1.30.67:3306)/go-test?charset=utf8"

var sqlxDB *sqlx.DB

type testModel struct {
	ID   string `db:"id,user"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (m testModel) GetID() string {
	return m.ID
}

func init() {
	sqlxDB, _ = sqlx.Open("mysql", connString)
}
