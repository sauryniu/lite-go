//go:generate mockgen -destination i-unit-of-work_mock.go -package db github.com/ahl5esoft/lite-go/db IUnitOfWork

package db

// IUnitOfWork is 工作单元
type IUnitOfWork interface {
	Commit() error
}
