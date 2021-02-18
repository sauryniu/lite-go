package db

// IQuery is 查询接口
type IQuery interface {
	Count() (int64, error)
	Order(fields ...string) IQuery
	OrderByDesc(fields ...string) IQuery
	Skip(v int) IQuery
	Take(v int) IQuery
	ToArray(dst interface{}) error
	Where(args ...interface{}) IQuery
}
