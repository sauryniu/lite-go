package contract

// IDbFactory is 数据工厂
type IDbFactory interface {
	Db(entry IDbModel, extra ...interface{}) IDbRepository
	Uow() IUnitOfWork
}
