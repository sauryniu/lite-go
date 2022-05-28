package contract

// 数据查询接口
type IDbQuery interface {
	// 查询数量
	Count() (int64, error)
	// 排序(正序)
	Order(fields ...string) IDbQuery
	// 排序(倒序)
	OrderByDesc(fields ...string) IDbQuery
	// 跳过行数
	Skip(v int) IDbQuery
	// 限制行数
	Take(v int) IDbQuery
	// 查询行数
	ToArray(dst interface{}) error
	// 条件
	Where(args ...interface{}) IDbQuery
}
