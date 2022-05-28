package mongosvc

import (
	"reflect"
)

var (
	testDriverFactory = newDriverFactory("lite-go-test", "mongodb://localhost:27017")
	testModelMetadata = getModelMetadata(testModelType)
	testModelType     = reflect.TypeOf(testModel{})
)

type testModel struct {
	ID  string `alias:"user" bson:"_id" db:"_id"`
	Int int
	Str string `db:"str"`
}

func (m testModel) GetID() string {
	return m.ID
}
