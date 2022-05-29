package dbop

// 表操作值
type Value string

const (
	// 增加
	Insert Value = "c"
	// 删除
	Delete Value = "d"
	// 更新
	Update Value = "u"
	// 查询
	Query Value = "r"
)
