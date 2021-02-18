package mongodb

import (
	"github.com/ahl5esoft/lite-go/db"
	"github.com/ahl5esoft/lite-go/db/identity"
)

type repository struct {
	isTx        bool
	modelStruct identity.IStruct
	pool        *connectPool
	uow         *unitOfWork
}

func (m repository) Add(entry identity.IIdentity) error {
	m.uow.registerAdd(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m repository) Query() db.IQuery {
	return newQuery(m.pool, m.modelStruct)
}

func (m repository) Remove(entry identity.IIdentity) error {
	m.uow.registerRemove(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m repository) Save(entry identity.IIdentity) error {
	m.uow.registerSave(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func newRepository(pool *connectPool, modelStruct identity.IStruct, uow *unitOfWork, isTx bool) db.IRepository {
	return &repository{
		isTx:        isTx,
		modelStruct: modelStruct,
		pool:        pool,
		uow:         uow,
	}
}
