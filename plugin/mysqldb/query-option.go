package mysqldb

type orderOption struct {
	Field string
	Flag  string
}

type queryOption struct {
	Orders    []orderOption
	Skip      int
	Take      int
	Where     string
	WhereArgs []interface{}
}
