package contract

// 工作单元仓储
type IUnitOfWorkRepository interface {
	// 工作单元
	IUnitOfWork

	// 注册删除
	RegisterDelete(entry IDbModel)
	// 注册新增
	RegisterInsert(entry IDbModel)
	// 注册更新
	RegisterUpdate(entry IDbModel)
}
