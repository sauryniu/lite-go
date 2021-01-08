package mysqldb

import (
	"fmt"
	"reflect"
	"testing"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/stretchr/testify/assert"
)

func Test_query_Count(t *testing.T) {
	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(testModel{}),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	res, err := self.Count()
	assert.Nil(t, err)
	assert.Equal(t, res, int64(0))
}

func Test_query_Count_有数据(t *testing.T) {
	uow := &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
	var entries []testModel
	underscore.Range(0, 25, 1).Map(func(r int, _ int) testModel {
		entry := testModel{
			ID:   fmt.Sprintf("id-%d", r),
			Name: fmt.Sprintf("name-%d", r),
			Age:  10 + r,
		}
		uow.RegisterAdd(entry)
		return entry
	}).Value(&entries)

	uow.Commit()

	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entries[0]),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	res, err := self.Count()

	underscore.Chain(entries).Each(func(r testModel, _ int) {
		uow.RegisterRemove(r)
	})
	uow.Commit()

	assert.Nil(t, err)
	assert.Equal(
		t,
		res,
		int64(
			len(entries),
		),
	)
}

func Test_query_Count_有条件(t *testing.T) {
	uow := &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
	var entries []testModel
	underscore.Range(0, 4, 1).Map(func(r int, _ int) testModel {
		entry := testModel{
			ID:   fmt.Sprintf("id-%d", r),
			Name: fmt.Sprintf("name-%d", r),
			Age:  10 + r,
		}
		uow.RegisterAdd(entry)
		return entry
	}).Value(&entries)

	uow.Commit()

	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entries[0]),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	res, err := self.Where("age % 5 = ?", 0).Count()

	underscore.Chain(entries).Each(func(r testModel, _ int) {
		uow.RegisterRemove(r)
	})
	uow.Commit()

	assert.Nil(t, err)
	assert.Equal(
		t,
		res,
		int64(1),
	)
}

func Test_query_Order(t *testing.T) {
	self := &query{
		DB: sqlxDB,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	self.Order("id", "name")
	assert.Equal(t, self.Option, queryOption{
		Orders: []orderOption{
			{
				Field: "id",
				Flag:  "ASC",
			},
			{
				Field: "name",
				Flag:  "ASC",
			},
		},
	})
}

func Test_query_OrderByDesc(t *testing.T) {
	self := &query{
		DB: sqlxDB,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	self.OrderByDesc("id", "name")
	assert.Equal(t, self.Option, queryOption{
		Orders: []orderOption{
			{
				Field: "id",
				Flag:  "DESC",
			},
			{
				Field: "name",
				Flag:  "DESC",
			},
		},
	})
}

func Test_query_Skip(t *testing.T) {
	self := &query{
		DB: sqlxDB,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	self.Skip(15)
	assert.Equal(t, self.Option, queryOption{
		Orders: make([]orderOption, 0),
		Skip:   15,
	})
}

func Test_query_Take(t *testing.T) {
	self := &query{
		DB: sqlxDB,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	self.Take(15)
	assert.Equal(t, self.Option, queryOption{
		Orders: make([]orderOption, 0),
		Take:   15,
	})
}

func Test_query_ToArray(t *testing.T) {
	uow := &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
	var entries []testModel
	underscore.Range(0, 4, 1).Map(func(r int, _ int) testModel {
		entry := testModel{
			ID:   fmt.Sprintf("id-%d", r),
			Name: fmt.Sprintf("name-%d", r),
			Age:  10 + r,
		}
		uow.RegisterAdd(entry)
		return entry
	}).Value(&entries)

	uow.Commit()

	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entries[0]),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	var res []testModel
	err := self.ToArray(&res)

	underscore.Chain(entries).Each(func(r testModel, _ int) {
		uow.RegisterRemove(r)
	})
	uow.Commit()

	assert.Nil(t, err)
	assert.EqualValues(t, res, entries)
}

func Test_query_ToArray_条件(t *testing.T) {
	uow := &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
	var entries []testModel
	underscore.Range(0, 4, 1).Map(func(r int, _ int) testModel {
		entry := testModel{
			ID:   fmt.Sprintf("id-%d", r),
			Name: fmt.Sprintf("name-%d", r),
			Age:  10 + r,
		}
		uow.RegisterAdd(entry)
		return entry
	}).Value(&entries)

	uow.Commit()

	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entries[0]),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	var res []testModel
	err := self.Where(
		fmt.Sprintf("age = %d", entries[0].Age),
	).ToArray(&res)

	underscore.Chain(entries).Each(func(r testModel, _ int) {
		uow.RegisterRemove(r)
	})
	uow.Commit()

	assert.Nil(t, err)
	assert.EqualValues(
		t,
		res,
		[]testModel{entries[0]},
	)
}

func Test_query_ToArray_排序(t *testing.T) {
	uow := &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
	var entries []testModel
	underscore.Range(0, 4, 1).Map(func(r int, _ int) testModel {
		entry := testModel{
			ID:   fmt.Sprintf("id-%d", r),
			Name: fmt.Sprintf("name-%d", r),
			Age:  10 + r,
		}
		uow.RegisterAdd(entry)
		return entry
	}).Value(&entries)

	uow.Commit()

	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entries[0]),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	var res []testModel
	err := self.OrderByDesc("age").ToArray(&res)

	underscore.Chain(entries).Each(func(r testModel, _ int) {
		uow.RegisterRemove(r)
	})
	uow.Commit()

	var sorted []testModel
	underscore.Range(3, -1, -1).Map(func(r int, _ int) testModel {
		return entries[r]
	}).Value(&sorted)

	assert.Nil(t, err)
	assert.EqualValues(t, res, sorted)
}

func Test_query_ToArray_限制(t *testing.T) {
	uow := &unitOfWork{
		DB:    sqlxDB,
		Items: make([]unitOfWorkItem, 0),
	}
	var entries []testModel
	underscore.Range(0, 4, 1).Map(func(r int, _ int) testModel {
		entry := testModel{
			ID:   fmt.Sprintf("id-%d", r),
			Name: fmt.Sprintf("name-%d", r),
			Age:  10 + r,
		}
		uow.RegisterAdd(entry)
		return entry
	}).Value(&entries)

	uow.Commit()

	self := &query{
		DB:        sqlxDB,
		ModelType: reflect.TypeOf(entries[0]),
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	var res []testModel
	err := self.Skip(1).Take(10).ToArray(&res)

	var actual []testModel
	underscore.Chain(entries).Where(func(r testModel, i int) bool {
		uow.RegisterRemove(r)

		return i > 0
	}).Value(&actual)
	uow.Commit()

	assert.Nil(t, err)
	assert.EqualValues(t, res, actual)
}

func Test_query_Where(t *testing.T) {
	self := &query{
		DB: sqlxDB,
		Option: queryOption{
			Orders: make([]orderOption, 0),
		},
	}
	self.Where("age > ?", 5)
	assert.Equal(t, self.Option, queryOption{
		Orders:    make([]orderOption, 0),
		Where:     "age > ?",
		WhereArgs: []interface{}{5},
	})
}
