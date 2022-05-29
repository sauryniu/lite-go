package dbsvc

import "github.com/ahl5esoft/lite-go/contract"

type dbRepository struct {
	uow             contract.IUnitOfWorkRepository
	isTx            bool
	createQueryFunc func() contract.IDbQuery
}

func (m dbRepository) Delete(entry contract.IDbModel) error {
	m.uow.RegisterDelete(entry)
	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m dbRepository) Insert(entry contract.IDbModel) error {
	m.uow.RegisterInsert(entry)
	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m dbRepository) Query() contract.IDbQuery {
	return m.createQueryFunc()
}

func (m dbRepository) Update(entry contract.IDbModel) error {
	m.uow.RegisterUpdate(entry)
	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

// 创建数据仓储
func NewDbRepository(
	uow contract.IUnitOfWorkRepository,
	isTx bool,
	createQueryFunc func() contract.IDbQuery,
) contract.IDbRepository {
	return &dbRepository{
		createQueryFunc: createQueryFunc,
		isTx:            isTx,
		uow:             uow,
	}
}
