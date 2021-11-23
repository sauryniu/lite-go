package mongosvc

import (
	"reflect"
)

const (
	testUri    = "mongodb://localhost:27017"
	testDbName = "infrastructure-core-test"
)

var (
	testClient        = newClientWrapper(testDbName, testUri)
	testModelMetadata = getModelMetadata(testModelType)
	testModelType     = reflect.TypeOf(testModel{})
)

type testModel struct {
	ID   string `alias:"user" bson:"_id" db:"_id"`
	Name string `db:"name"`
	Age  int
}

func (m testModel) GetID() string {
	return m.ID
}

type testTimeAssociateModel struct {
	ID         string `alias:"user" bson:"_id" db:"_id"`
	CreatedOn  int64
	DeletedOn  int64
	ModifiedOn int64
}

func (m testTimeAssociateModel) GetID() string {
	return m.ID
}

func (m *testTimeAssociateModel) SetCreatedOn(v int64) {
	m.CreatedOn = v
}

func (m *testTimeAssociateModel) SetDeletedOn(v int64) {
	m.DeletedOn = v
}

func (m *testTimeAssociateModel) SetModifiedOn(v int64) {
	m.ModifiedOn = v
}
