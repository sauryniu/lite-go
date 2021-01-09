package mysqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factory_Db(t *testing.T) {
	self, err := New(FactoryOption{
		DbName:   "go-test",
		Host:     "10.1.33.67",
		Password: "123456",
		Username: "root",
	})

	if err == nil {
		var res []testModel
		self.Db(testModel{}).Query().ToArray(&res)
		assert.Empty(t, res)
	}

	assert.Nil(t, err)
}

func Test_factory_Db_extra(t *testing.T) {
	self, err := New(FactoryOption{
		DbName:   "go-test",
		Host:     "10.1.33.67",
		Password: "123456",
		Username: "root",
	})

	if err == nil {
		uow := self.Uow()
		db := self.Db(testModel{}, uow)
		entry := testModel{
			ID:   "id-1",
			Name: "add",
			Age:  11,
		}
		db.Add(entry)

		var res1 []testModel
		db.Query().ToArray(&res1)

		uow.Commit()

		var res2 []testModel
		db.Query().ToArray(&res2)

		db.Remove(entry)
		uow.Commit()

		assert.Empty(t, res1)
		assert.EqualValues(
			t,
			res2,
			[]testModel{entry},
		)
	}

	assert.Nil(t, err)
}

func Test_factory_Uow(t *testing.T) {
	self, err := New(FactoryOption{
		DbName:   "go-test",
		Host:     "10.1.33.67",
		Password: "123456",
		Username: "root",
	})

	if err == nil {
		uow := self.Uow().(*unitOfWork)
		entry := testModel{
			ID:   "id-1",
			Name: "add",
			Age:  11,
		}
		uow.RegisterAdd(entry)
		entry.Name = "save"
		uow.RegisterSave(entry)
		uow.RegisterRemove(entry)
		assert.Nil(
			t,
			uow.Commit(),
		)
	}

	assert.Nil(t, err)
}
