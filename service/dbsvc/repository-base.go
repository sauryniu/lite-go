package dbsvc

import "github.com/ahl5esoft/lite-go/contract"

type repositoryBase struct {
	createQueryFunc func() contract.IDbQuery
	isTx            bool
	nowTime         contract.INowTime
	uow             contract.IUnitOfWorkRepository
}

func (m repositoryBase) Add(entry contract.IDbModel) error {
	if setter, ok := entry.(contract.IModelCreatedOnSetter); ok {
		setter.SetCreatedOn(
			m.nowTime.Unix(),
		)
	}

	m.uow.RegisterAdd(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m repositoryBase) Query() contract.IDbQuery {
	return m.createQueryFunc()
}

func (m repositoryBase) Remove(entry contract.IDbModel) error {
	if setter, ok := entry.(contract.IModelDeletedOnSetter); ok {
		setter.SetDeletedOn(
			m.nowTime.Unix(),
		)
		m.uow.RegisterSave(entry)
	} else {
		m.uow.RegisterRemove(entry)
	}

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func (m repositoryBase) Save(entry contract.IDbModel) error {
	if setter, ok := entry.(contract.IModelModifiedOnSetter); ok {
		setter.SetModifiedOn(
			m.nowTime.Unix(),
		)
	}

	m.uow.RegisterSave(entry)

	if m.isTx {
		return nil
	}

	return m.uow.Commit()
}

func NewRepositoryBase(
	createQueryFunc func() contract.IDbQuery,
	isTx bool,
	nowTime contract.INowTime,
	uow contract.IUnitOfWorkRepository,
) contract.IDbRepository {
	return &repositoryBase{
		createQueryFunc: createQueryFunc,
		isTx:            isTx,
		nowTime:         nowTime,
		uow:             uow,
	}
}
