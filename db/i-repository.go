package db

import "github.com/ahl5esoft/lite-go/db/identity"

// IRepository is 表仓库
type IRepository interface {
	Add(entry identity.IIdentity) error
	Query() IQuery
	Remove(entry identity.IIdentity) error
	Save(entry identity.IIdentity) error
}
