package dbsvc

import "github.com/ahl5esoft/lite-go/contract"

type repository struct {
	uow             contract.IUnitOfWorkRepository
	isTx            bool
	createQueryFunc func() contract.IDbQuery
}

func (m repository) Delete(entry contract.IDbModel) error {
	m.uow.RegisterDelete(entry)
	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m repository) Insert(entry contract.IDbModel) error {
	m.uow.RegisterInsert(entry)
	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m repository) Query() contract.IDbQuery {
	return m.createQueryFunc()
}

func (m repository) Update(entry contract.IDbModel) error {
	m.uow.RegisterUpdate(entry)
	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

// 创建数据仓储
func NewRepository(
	uow contract.IUnitOfWorkRepository,
	isTx bool,
	createQueryFunc func() contract.IDbQuery,
) contract.IDbRepository {
	return &repository{
		createQueryFunc: createQueryFunc,
		isTx:            isTx,
		uow:             uow,
	}
}
