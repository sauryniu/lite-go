package mongodb

import (
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
)

type repository struct {
	IsTx   bool
	Pool   *connectPool
	Struct identity.IStruct
	Uow    *unitOfWork
}

func (m repository) Add(entry identity.IIdentity) error {
	m.Uow.RegisterAdd(entry)

	if m.IsTx {
		return nil
	}

	return m.Uow.Commit()
}

func (m repository) Query() db.IQuery {
	return newQuery(m.Pool, m.Struct)
}

func (m repository) Remove(entry identity.IIdentity) error {
	m.Uow.RegisterRemove(entry)

	if m.IsTx {
		return nil
	}

	return m.Uow.Commit()
}

func (m repository) Save(entry identity.IIdentity) error {
	m.Uow.RegisterSave(entry)

	if m.IsTx {
		return nil
	}

	return m.Uow.Commit()
}

func newRepository(pool *connectPool, s identity.IStruct, uow *unitOfWork, isTx bool) db.IRepository {
	return &repository{
		IsTx:   isTx,
		Pool:   pool,
		Struct: s,
		Uow:    uow,
	}
}
