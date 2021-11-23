package dbsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testModel struct {
	ID string
}

func (m testModel) GetID() string {
	return m.ID
}

type testCreatedOnModel struct {
	CreatedOn int64
	ID        string
}

func (m testCreatedOnModel) GetID() string {
	return m.ID
}

func (m *testCreatedOnModel) SetCreatedOn(v int64) {
	m.CreatedOn = v
}

type testDeletedOnModel struct {
	DeletedOn int64
	ID        string
}

func (m testDeletedOnModel) GetID() string {
	return m.ID
}

func (m *testDeletedOnModel) SetDeletedOn(v int64) {
	m.DeletedOn = v
}

type testModifiedOnOnModel struct {
	ModifiedOn int64
	ID         string
}

func (m testModifiedOnOnModel) GetID() string {
	return m.ID
}

func (m *testModifiedOnOnModel) SetModifiedOn(v int64) {
	m.ModifiedOn = v
}

func Test_repositoryBase_Add(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterAdd(entry)

		mockUow.EXPECT().Commit().Return(nil)

		err := self.Add(entry)
		assert.NoError(t, err)
	})

	t.Run("事务", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)
		self.isTx = true

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterAdd(entry)

		err := self.Add(entry)
		assert.NoError(t, err)
	})

	t.Run("CreatedOn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)
		self.isTx = true

		mockNowTime := contract.NewMockINowTime(ctrl)
		self.nowTime = mockNowTime

		unix := int64(1)
		mockNowTime.EXPECT().Unix().Return(unix)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testCreatedOnModel{}
		mockUow.EXPECT().RegisterAdd(&entry)

		err := self.Add(&entry)
		assert.NoError(t, err)
		assert.Equal(t, entry.CreatedOn, unix)
	})
}

func Test_repositoryBase_Remove(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterRemove(entry)

		mockUow.EXPECT().Commit().Return(nil)

		err := self.Remove(entry)
		assert.NoError(t, err)
	})

	t.Run("事务", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)
		self.isTx = true

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterRemove(entry)

		err := self.Remove(entry)
		assert.NoError(t, err)
	})

	t.Run("DeletedOn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)
		self.isTx = true

		mockNowTime := contract.NewMockINowTime(ctrl)
		self.nowTime = mockNowTime

		unix := int64(1)
		mockNowTime.EXPECT().Unix().Return(unix)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testDeletedOnModel{}
		mockUow.EXPECT().RegisterSave(&entry)

		err := self.Remove(&entry)
		assert.NoError(t, err)
		assert.Equal(t, entry.DeletedOn, unix)
	})
}

func Test_repositoryBase_Save(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterSave(entry)

		mockUow.EXPECT().Commit().Return(nil)

		err := self.Save(entry)
		assert.NoError(t, err)
	})

	t.Run("事务", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)
		self.isTx = true

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModel{}
		mockUow.EXPECT().RegisterSave(entry)

		err := self.Save(entry)
		assert.NoError(t, err)
	})

	t.Run("ModifiedOn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(repositoryBase)
		self.isTx = true

		mockNowTime := contract.NewMockINowTime(ctrl)
		self.nowTime = mockNowTime

		unix := int64(1)
		mockNowTime.EXPECT().Unix().Return(unix)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		self.uow = mockUow

		entry := testModifiedOnOnModel{}
		mockUow.EXPECT().RegisterSave(&entry)

		err := self.Save(&entry)
		assert.NoError(t, err)
		assert.Equal(t, entry.ModifiedOn, unix)
	})
}
