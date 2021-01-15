package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/mongodb"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type mongoAPI struct {
	DbFactory db.IFactory `inject:"db"`
}

func (m mongoAPI) Call() (interface{}, error) {
	uow := m.DbFactory.Uow()
	db := m.DbFactory.Db(mongoModel{}, uow)
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
}

type mongoModel struct {
	ID  string `db:"_id,go-test" bson:"_id"`
	Age int
}

func (m mongoModel) GetID() string {
	return m.ID
}

func Test_GinMongoAPI(t *testing.T) {
	dbFactory, err := mongodb.New(mongodb.FactoryOption{
		DbName: "example",
		URI:    "mongodb://localhost:27017",
	})
	assert.NoError(t, err)
	ioc.Set(db.IoCKey, dbFactory)

	endpoint := "endpoint"
	name := "mongo"
	api.Register(endpoint, name, mongoAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	resp := httptest.NewRecorder()
	ginex.Run(
		gin.ReleaseMode,
		ginex.NewPostRunOption(),
		ginex.NewServeHTTPRunOption(req, resp),
	)

	var entries []mongoModel
	underscore.Range(10, 15, 1).Map(func(r int, _ int) mongoModel {
		return mongoModel{
			ID:  fmt.Sprintf("id-%d", r),
			Age: r,
		}
	}).Value(&entries)
	res, _ := jsoniter.MarshalToString(api.Response{
		Data:  entries,
		Error: 0,
	})
	assert.JSONEq(
		t,
		resp.Body.String(),
		res,
	)
}
