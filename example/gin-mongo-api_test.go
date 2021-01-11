package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/mongodb"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type ginMongoAPIContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *ginMongoAPIContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "mongo"
	api.Register(endpoint, name, func() api.IAPI {
		var a *mongoAPI
		a = &mongoAPI{
			API: &ginex.API{
				CallFunc: func() (interface{}, error) {
					uow := a.DbFactory.Uow()
					db := a.DbFactory.Db(mongoModel{}, uow)
					underscore.Range(10, 15, 1).Each(func(r int, _ int) {
						db.Add(mongoModel{
							ID:  fmt.Sprintf("id-%d", r),
							Age: r,
						})
					})
					if err := uow.Commit(); err != nil {
						return nil, err
					}

					var entries []mongoModel
					if err := db.Query().ToArray(&entries); err != nil {
						return nil, err
					}

					underscore.Chain(entries).Each(func(r mongoModel, _ int) {
						db.Remove(r)
					})
					if err := uow.Commit(); err != nil {
						return nil, err
					}

					return entries, nil
				},
			},
		}
		return a
	})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m ginMongoAPIContext) GetGinMode() string {
	return ""
}

func (m ginMongoAPIContext) GetGinPort() int {
	return 0
}

func (m ginMongoAPIContext) GetMongoOption() mongodb.FactoryOption {
	return mongodb.FactoryOption{
		DbName: "example",
		URI:    "mongodb://localhost:27017",
	}
}

type mongoAPI struct {
	*ginex.API

	DbFactory db.IFactory `inject:"db"`
}

type mongoModel struct {
	ID  string `db:"_id,go-test" bson:"_id"`
	Age int
}

func (m mongoModel) GetID() string {
	return m.ID
}

func Test_GinMongoAPI(t *testing.T) {
	handler := mongodb.NewStartupHandler()
	handler.SetNext(
		ginex.NewStartupHandler(),
	)
	ctx := new(ginMongoAPIContext)
	err := handler.Handle(ctx)
	assert.NoError(t, err)

	var entries []mongoModel
	underscore.Range(10, 15, 1).Map(func(r int, _ int) mongoModel {
		return mongoModel{
			ID:  fmt.Sprintf("id-%d", r),
			Age: r,
		}
	}).Value(&entries)
	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": entries,
		"err":  0,
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
