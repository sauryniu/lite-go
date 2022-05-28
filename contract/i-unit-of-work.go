package contract

// 工作单元
type IUnitOfWork interface {
	// 提交
	Commit() error
}
