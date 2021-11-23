package contract

// IDbRepository is 数据仓库
type IDbRepository interface {
	Add(entry IDbModel) error
	Query() IDbQuery
	Remove(entry IDbModel) error
	Save(entry IDbModel) error
}
