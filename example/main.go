package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/api/bg"
	dbop "github.com/ahl5esoft/lite-go/example/admin-permission/model/enum/db-op"
	"github.com/ahl5esoft/lite-go/example/admin-permission/model/global"
	"github.com/ahl5esoft/lite-go/service/mongosvc"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	gin.SetMode(gin.TestMode)
	app := gin.New()

	dbName := "lite-go-admin-permission"
	mongoUri := "mongodb://localhost:27017"
	dbFactory := mongosvc.NewDbFactory(dbName, mongoUri)

	defer func() {
		ctx := context.Background()
		option := options.Client().ApplyURI(mongoUri)
		client, _ := mongo.Connect(ctx, option)
		client.Database(dbName).Drop(ctx)
	}()

	uow := dbFactory.Uow()

	adminPermissionDb := dbFactory.Db(global.AdminPermission{}, uow)
	bsonBf, _ := bson.Marshal(bson.M{
		"_id": bson.M{
			"$in": []string{"admin-1", "admin-2"},
		},
	})
	adminPermissionDb.Insert(global.AdminPermission{
		ID: "admin-1",
		Permission: map[string]map[dbop.Value]interface{}{
			"Admin": {
				dbop.Query: bsonBf,
			},
		},
	})
	adminPermissionDb.Insert(global.AdminPermission{
		ID: "admin-2",
		Permission: map[string]map[dbop.Value]interface{}{
			"Admin": {
				dbop.Delete: true,
				dbop.Insert: true,
				dbop.Update: true,
			},
		},
	})

	adminDb := dbFactory.Db(global.Admin{}, uow)
	for i := 1; i <= 5; i++ {
		adminDb.Insert(global.Admin{
			ID:   fmt.Sprintf("admin-%d", i),
			Name: fmt.Sprintf("name-%d", i),
		})
	}

	uow.Commit()

	app.POST("/r", func(ctx *gin.Context) {
		ctx.Set("admin", "admin-1")

		var api contract.IApi
		api = new(bg.QueryApi)
		api.(*bg.QueryApi).MongoDbFactory = dbFactory
		api.(contract.IApiSession).InitSession(ctx)
		res, err := api.Call()
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": res,
			"err":  err,
		})
	})

	for _, r := range []dbop.Value{dbop.Delete, dbop.Insert, dbop.Update} {
		(func(op dbop.Value) {
			app.POST(
				fmt.Sprintf("/%s", op),
				func(ctx *gin.Context) {
					ctx.Set("admin", "admin-2")

					var api contract.IApi
					api = new(bg.CUDApi)
					api.(*bg.CUDApi).MongoDbFactory = dbFactory
					api.(*bg.CUDApi).DbOp = op
					api.(contract.IApiSession).InitSession(ctx)
					res, err := api.Call()
					if err == nil {
						err = fmt.Errorf("")
					}
					ctx.JSON(http.StatusOK, map[string]interface{}{
						"data": res,
						"err":  err.Error(),
					})
				},
			)
		})(r)
	}

	for _, r := range []dbop.Value{dbop.Delete, dbop.Query, dbop.Insert, dbop.Update} {
		req := httptest.NewRequest(
			"POST",
			fmt.Sprintf("/%s", r),
			strings.NewReader(``),
		)
		resp := httptest.NewRecorder()
		app.ServeHTTP(resp, req)
		res, err := ioutil.ReadAll(
			resp.Result().Body,
		)
		fmt.Println(
			string(res),
			err,
		)
	}
}
