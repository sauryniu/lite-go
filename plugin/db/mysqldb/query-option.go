package mysqldb

type queryOption struct {
	Orders    []orderOption
	Skip      int
	Take      int
	Where     string
	WhereArgs []interface{}
}
