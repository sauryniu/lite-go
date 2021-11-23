package contract

type IUnitOfWork interface {
	Commit() error
}
