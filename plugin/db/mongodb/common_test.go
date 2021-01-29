package mongodb

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/plugin/db/identity"
)

const (
	uri    = "mongodb://localhost:27017"
	dbName = "lite-go"
)

var (
	pool       = newPool(uri, dbName)
	testStruct = identity.NewStruct(
		reflect.TypeOf(testModel{}),
	)
)

type testModel struct {
	ID   string `alias:"user" bson:"_id" db:"_id"`
	Name string `db:"name"`
	Age  int
}

func (m testModel) GetID() string {
	return m.ID
}
