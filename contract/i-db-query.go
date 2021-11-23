package contract

// IDbQuery is 数据查询接口
type IDbQuery interface {
	Count() (int64, error)
	Order(fields ...string) IDbQuery
	OrderByDesc(fields ...string) IDbQuery
	Skip(v int) IDbQuery
	Take(v int) IDbQuery
	ToArray(dst interface{}) error
	Where(args ...interface{}) IDbQuery
}
