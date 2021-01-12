package ginapi

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
	"github.com/ahl5esoft/lite-go/plugin/redisex/goredis"
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

func (m mongoAPI) GetScope() int {
	return 0
}

type mongoModel struct {
	ID  string `db:"_id,go-test" bson:"_id"`
	Age int
}

func (m mongoModel) GetID() string {
	return m.ID
}

type startupMongoContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *startupMongoContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "mongo"
	api.Register(endpoint, name, mongoAPI{})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m startupMongoContext) GetGinMode() string {
	return ""
}

func (m startupMongoContext) GetGinPort() int {
	return 0
}

func (m startupMongoContext) GetMongoOption() mongodb.FactoryOption {
	return mongodb.FactoryOption{
		DbName: "example",
		URI:    "mongodb://localhost:27017",
	}
}

func (m startupMongoContext) HandleGinCtx(ctx *gin.Context) {
	handleGinContet(ctx)
}

func Test_GinMongoAPI(t *testing.T) {
	ctx := new(startupMongoContext)
	handler := mongodb.NewStartupHandler()
	handler.SetNext(
		goredis.NewStartupHandler(),
	).SetNext(
		ginex.NewStartupHandler(),
	)
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
