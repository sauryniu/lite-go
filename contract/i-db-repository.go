package contract

// 数据仓库
type IDbRepository interface {
	// 删除
	Delete(entry IDbModel) error
	// 新增
	Insert(entry IDbModel) error
	// 创建查询
	Query() IDbQuery
	// 更新
	Update(entry IDbModel) error
}
