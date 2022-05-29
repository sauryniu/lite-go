package dbsvc

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_query_Count(t *testing.T) {
	t.Run("无限制", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		model := "Admin"
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newDbQuery(mockOriginDbQuery, model, dbFactory)

		dbFactory.adminPermissions = make([]global.AdminPermission, 0)

		mockOriginDbQuery.EXPECT().Where(bson.M{}).Return(mockOriginDbQuery)

		mockOriginDbQuery.EXPECT().Count().Return(
			int64(1),
			nil,
		)

		res, err := self.Count()
		assert.NoError(t, err)
		assert.Equal(
			t,
			res,
			int64(1),
		)
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		model := "Admin"
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newDbQuery(mockOriginDbQuery, model, dbFactory)

		dbFactory.adminPermissions = []global.AdminPermission{
			{
				Permission: map[string]map[dbop.Value]interface{}{
					"Admin": {
						dbop.Query: bson.M{
							"name": "admin",
						},
					},
				},
			},
		}

		mockOriginDbQuery.EXPECT().Where(bson.M{
			"name": "admin",
		}).Return(mockOriginDbQuery)

		mockOriginDbQuery.EXPECT().Count().Return(
			int64(1),
			nil,
		)

		res, err := self.Count()
		assert.NoError(t, err)
		assert.Equal(
			t,
			res,
			int64(1),
		)
	})
}

func Test_query_Order(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		self := newDbQuery(mockOriginDbQuery, "", nil)

		mockOriginDbQuery.EXPECT().Order("a", "b")

		self.Order("a", "b")
	})
}

func Test_query_OrderByDesc(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		self := newDbQuery(mockOriginDbQuery, "", nil)

		mockOriginDbQuery.EXPECT().OrderByDesc("a", "b")

		self.OrderByDesc("a", "b")
	})
}

func Test_query_Skip(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		self := newDbQuery(mockOriginDbQuery, "", nil)

		mockOriginDbQuery.EXPECT().Skip(10)

		self.Skip(10)
	})
}

func Test_query_Take(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		self := newDbQuery(mockOriginDbQuery, "", nil)

		mockOriginDbQuery.EXPECT().Take(20)

		self.Take(20)
	})
}

func Test_query_ToArray(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockOriginDbQuery := contract.NewMockIDbQuery(ctrl)
		model := "Admin"
		dbFactory := NewDbFactory(nil, "").(*dbFactory)
		self := newDbQuery(mockOriginDbQuery, model, dbFactory)

		dbFactory.adminPermissions = []global.AdminPermission{
			{
				Permission: map[string]map[dbop.Value]interface{}{
					"Admin": {
						dbop.Query: bson.M{
							"name": "admin",
						},
					},
				},
			},
		}

		mockOriginDbQuery.EXPECT().Where(bson.M{
			"name": "admin",
		}).Return(mockOriginDbQuery)

		mockOriginDbQuery.EXPECT().ToArray(
			gomock.Any(),
		).SetArg(0, []global.Admin{
			{},
			{},
		}).Return(nil)

		var res []global.Admin
		err := self.ToArray(&res)
		assert.NoError(t, err)
		assert.EqualValues(t, res, []global.Admin{
			{},
			{},
		})
	})
}

func Test_query_Where(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := newDbQuery(nil, "", nil)

		filter := bson.M{
			"name": "",
		}
		self.Where(filter)
		assert.EqualValues(
			t,
			self.(*dbQuery).filter,
			filter,
		)
	})
}
