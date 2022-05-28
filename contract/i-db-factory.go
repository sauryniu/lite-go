package contract

// 数据工厂
type IDbFactory interface {
	// 创建数据仓储
	Db(entry IDbModel, extra ...interface{}) IDbRepository
	// 创建工作单元
	Uow() IUnitOfWork
}
