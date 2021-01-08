package mongodb

import (
	"reflect"

	"github.com/ahl5esoft/lite-go/plugin/db/identity"
)

type testModel struct {
	ID   string `db:"_id,user" bson:"_id"`
	Name string `db:"name"`
	Age  int
}

func (m testModel) GetID() string {
	return m.ID
}

var (
	pool       = newPool("lite-go", "mongodb://localhost:27017")
	testStruct = identity.NewStruct(
		reflect.TypeOf(testModel{}),
	)
)
