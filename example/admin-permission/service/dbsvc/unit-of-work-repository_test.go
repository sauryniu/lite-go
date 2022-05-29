package dbsvc

import (
	"fmt"
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_unitOfWorkRepository_Commit(t *testing.T) {
	t.Run("无权限限制", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newUnitOfWorkRepository(mockUow, dbFactory)

		self.(*unitOfWorkRepository).modelOperation = map[string]map[dbop.Value]bool{
			"Admin": {
				dbop.Delete: true,
				dbop.Insert: true,
				dbop.Update: true,
			},
		}

		dbFactory.adminPermissions = make([]global.AdminPermission, 0)

		mockUow.EXPECT().Commit()

		err := self.Commit()
		assert.NoError(t, err)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			make(map[string]map[dbop.Value]bool),
		)
	})

	t.Run("限制删除", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newUnitOfWorkRepository(mockUow, dbFactory)

		self.(*unitOfWorkRepository).modelOperation = map[string]map[dbop.Value]bool{
			"Admin": {
				dbop.Delete: true,
				dbop.Insert: true,
				dbop.Update: true,
			},
		}

		dbFactory.adminPermissions = []global.AdminPermission{
			{
				Permission: map[string]map[dbop.Value]interface{}{
					"Admin": {
						dbop.Delete: true,
					},
				},
			},
		}

		err := self.Commit()
		assert.Error(t, err)
		assert.Equal(
			t,
			err.Error(),
			fmt.Sprintf("无权限: Admin, %s", dbop.Delete),
		)
	})

	t.Run("限制新增", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newUnitOfWorkRepository(mockUow, dbFactory)

		self.(*unitOfWorkRepository).modelOperation = map[string]map[dbop.Value]bool{
			"Admin": {
				dbop.Delete: true,
				dbop.Insert: true,
				dbop.Update: true,
			},
		}

		dbFactory.adminPermissions = []global.AdminPermission{
			{
				Permission: map[string]map[dbop.Value]interface{}{
					"Admin": {
						dbop.Insert: true,
					},
				},
			},
		}

		err := self.Commit()
		assert.Error(t, err)
		assert.Equal(
			t,
			err.Error(),
			fmt.Sprintf("无权限: Admin, %s", dbop.Insert),
		)
	})

	t.Run("限制更新", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newUnitOfWorkRepository(mockUow, dbFactory)

		self.(*unitOfWorkRepository).modelOperation = map[string]map[dbop.Value]bool{
			"Admin": {
				dbop.Delete: true,
				dbop.Insert: true,
				dbop.Update: true,
			},
		}

		dbFactory.adminPermissions = []global.AdminPermission{
			{
				Permission: map[string]map[dbop.Value]interface{}{
					"Admin": {
						dbop.Update: true,
					},
				},
			},
		}

		err := self.Commit()
		assert.Error(t, err)
		assert.Equal(
			t,
			err.Error(),
			fmt.Sprintf("无权限: Admin, %s", dbop.Update),
		)
	})
}

func Test_unitOfWorkRepository_RegisterDelete(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self := newUnitOfWorkRepository(mockUow, nil)

		entry := global.Admin{
			ID: "delete",
		}
		mockUow.EXPECT().RegisterDelete(entry)

		self.RegisterDelete(entry)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			map[string]map[dbop.Value]bool{
				"Admin": {
					dbop.Delete: true,
				},
			},
		)
	})

	t.Run("multi", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self := newUnitOfWorkRepository(mockUow, nil)

		entry := global.Admin{
			ID: "delete",
		}
		mockUow.EXPECT().RegisterDelete(entry)
		mockUow.EXPECT().RegisterDelete(entry)

		self.RegisterDelete(entry)
		self.RegisterDelete(entry)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			map[string]map[dbop.Value]bool{
				"Admin": {
					dbop.Delete: true,
				},
			},
		)
	})
}

func Test_unitOfWorkRepository_RegisterInsert(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self := newUnitOfWorkRepository(mockUow, nil)

		entry := global.Admin{
			ID: "insert",
		}
		mockUow.EXPECT().RegisterInsert(entry)

		self.RegisterInsert(entry)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			map[string]map[dbop.Value]bool{
				"Admin": {
					dbop.Insert: true,
				},
			},
		)
	})

	t.Run("multi", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self := newUnitOfWorkRepository(mockUow, nil)

		entry := global.Admin{
			ID: "insert",
		}
		mockUow.EXPECT().RegisterInsert(entry)
		mockUow.EXPECT().RegisterInsert(entry)

		self.RegisterInsert(entry)
		self.RegisterInsert(entry)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			map[string]map[dbop.Value]bool{
				"Admin": {
					dbop.Insert: true,
				},
			},
		)
	})
}

func Test_unitOfWorkRepository_RegisterUpdate(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self := newUnitOfWorkRepository(mockUow, nil)

		entry := global.Admin{
			ID: "update",
		}
		mockUow.EXPECT().RegisterUpdate(entry)

		self.RegisterUpdate(entry)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			map[string]map[dbop.Value]bool{
				"Admin": {
					dbop.Update: true,
				},
			},
		)
	})

	t.Run("multi", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self := newUnitOfWorkRepository(mockUow, nil)

		entry := global.Admin{
			ID: "update",
		}
		mockUow.EXPECT().RegisterUpdate(entry)
		mockUow.EXPECT().RegisterUpdate(entry)

		self.RegisterUpdate(entry)
		self.RegisterUpdate(entry)

		assert.EqualValues(
			t,
			self.(*unitOfWorkRepository).modelOperation,
			map[string]map[dbop.Value]bool{
				"Admin": {
					dbop.Update: true,
				},
			},
		)
	})
}
