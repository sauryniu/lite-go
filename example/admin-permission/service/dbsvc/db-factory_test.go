package dbsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_dbFactory_GetAdminPermissions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewDbFactory(mockDbFactory, "").(*dbFactory)

		mockDb := contract.NewMockIDbRepository(ctrl)
		mockDbFactory.EXPECT().Db(global.AdminPermission{}).Return(mockDb)

		mockDbQuery := contract.NewMockIDbQuery(ctrl)
		mockDb.EXPECT().Query().Return(mockDbQuery)

		self.adminID = "admin-id"
		mockDbQuery.EXPECT().Where(bson.M{
			"_id": self.adminID,
		}).Return(mockDbQuery)

		mockDbQuery.EXPECT().ToArray(
			gomock.Any(),
		).SetArg(0, []global.AdminPermission{
			{},
			{},
		}).Return(nil)

		res, err := self.GetAdminPermissions()
		assert.NoError(t, err)
		assert.EqualValues(t, res, []global.AdminPermission{
			{},
			{},
		})
	})
}

func Test_dbFactory_Uow(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbFactory := contract.NewMockIDbFactory(ctrl)
		self := NewDbFactory(mockDbFactory, "").(*dbFactory)

		mockUow := contract.NewMockIUnitOfWorkRepository(ctrl)
		mockDbFactory.EXPECT().Uow().Return(mockUow)

		res := self.Uow()
		_, ok := res.(*unitOfWorkRepository)
		assert.True(t, ok)
		assert.Equal(
			t,
			res.(*unitOfWorkRepository).uow,
			mockUow,
		)
	})
}
