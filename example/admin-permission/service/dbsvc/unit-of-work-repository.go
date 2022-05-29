package dbsvc

import (
	"fmt"
	"reflect"

	"github.com/ahl5esoft/lite-go/contract"
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"
)

type unitOfWorkRepository struct {
	uow            contract.IUnitOfWorkRepository
	adminID        string
	modelOperation map[string]map[dbop.Value]bool
	dbFactory      *dbFactory
}

func (m *unitOfWorkRepository) Commit() (err error) {
	defer func() {
		m.modelOperation = make(map[string]map[dbop.Value]bool)
	}()

	var adminPermissions []global.AdminPermission
	adminPermissions, err = m.dbFactory.GetAdminPermissions()
	if err != nil {
		return
	}

	if len(adminPermissions) > 0 {
		for k, v := range m.modelOperation {
			for ck := range v {
				if sv, ok := adminPermissions[0].Permission[k]; ok {
					if ov, ok := sv[ck]; ok && ov.(bool) {
						return fmt.Errorf("无权限: %s(%s)", k, ck)
					}
				}
			}
		}
	}

	err = m.uow.Commit()
	return
}

func (m *unitOfWorkRepository) RegisterDelete(entry contract.IDbModel) {
	m.addModelOperation(entry, dbop.Delete)
	m.uow.RegisterDelete(entry)
}

func (m *unitOfWorkRepository) RegisterInsert(entry contract.IDbModel) {
	m.addModelOperation(entry, dbop.Insert)
	m.uow.RegisterInsert(entry)
}

func (m *unitOfWorkRepository) RegisterUpdate(entry contract.IDbModel) {
	m.addModelOperation(entry, dbop.Update)
	m.uow.RegisterUpdate(entry)
}

func (m *unitOfWorkRepository) addModelOperation(entry contract.IDbModel, op dbop.Value) {
	modelType := reflect.TypeOf(entry)
	model := modelType.Name()
	if _, ok := m.modelOperation[model]; !ok {
		m.modelOperation[model] = make(map[dbop.Value]bool)
	}

	m.modelOperation[model][op] = true
}

func newUnitOfWorkRepository(
	uow contract.IUnitOfWorkRepository,
	dbFactory *dbFactory,
) contract.IUnitOfWorkRepository {
	return &unitOfWorkRepository{
		dbFactory:      dbFactory,
		modelOperation: make(map[string]map[dbop.Value]bool),
		uow:            uow,
	}
}
