package mysqldb

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB        *sqlx.DB
	IsTx      bool
	ModelType reflect.Type
	Uow       *unitOfWork
}

func (m repository) Add(entry identity.IIdentity) error {
	if err := m.Uow.RegisterAdd(entry); err != nil || m.IsTx {
		return err
	}

	return m.Uow.Commit()
}

func (m repository) Query() db.IQuery {
	return &query{
		DB:        m.DB,
		ModelType: m.ModelType,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
}

func (m repository) Remove(entry identity.IIdentity) error {
	if err := m.Uow.RegisterRemove(entry); err != nil || m.IsTx {
		return err
	}

	return m.Uow.Commit()
}

func (m repository) Save(entry identity.IIdentity) error {
	if err := m.Uow.RegisterSave(entry); err != nil || m.IsTx {
		return err
	}

	return m.Uow.Commit()
}
