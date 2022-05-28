package dbop

// 表操作值
type Value int

const (
	// 增加
	Insert Value = iota
	// 删除
	Delete
	// 更新
	Update
	// 查询
	Query
)
