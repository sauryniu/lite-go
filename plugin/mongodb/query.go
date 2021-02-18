package mongodb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/db"
	"github.com/ahl5esoft/lite-go/db/identity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type query struct {
	filter      bson.M
	findOption  *options.FindOptions
	modelStruct identity.IStruct
	pool        *connectPool
	sorts       bson.D
}

func (m query) Count() (int64, error) {
	c, err := m.pool.GetCollection(m.modelStruct)
	if err != nil {
		return 0, err
	}

	defer m.Reset()
	return c.CountDocuments(m.pool.ctx, m.filter)
}

func (m *query) Order(fields ...string) db.IQuery {
	m.sort(1, fields)
	return m
}

func (m *query) OrderByDesc(fields ...string) db.IQuery {
	m.sort(-1, fields)
	return m
}

func (m *query) Reset() {
	m.filter = make(bson.M)
	m.findOption = options.Find()
	m.sorts = bson.D{}
}

func (m *query) Skip(v int) db.IQuery {
	m.findOption = m.findOption.SetSkip(
		int64(v),
	)
	return m
}

func (m *query) Take(v int) db.IQuery {
	m.findOption = m.findOption.SetLimit(
		int64(v),
	)
	return m
}

func (m *query) ToArray(dst interface{}) error {
	c, err := m.pool.GetCollection(m.modelStruct)
	if err != nil {
		return err
	}

	defer m.Reset()

	if len(m.sorts) > 0 {
		m.findOption = m.findOption.SetSort(m.sorts)
	}

	cur, err := c.Find(m.pool.ctx, m.filter, m.findOption)
	if err != nil {
		return err
	}

	sliceType := reflect.SliceOf(
		m.modelStruct.GetType(),
	)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	for cur.Next(m.pool.ctx) {
		value := reflect.New(
			m.modelStruct.GetType(),
		)
		temp := value.Interface()
		err = cur.Decode(temp)

		sliceValue = reflect.Append(
			sliceValue,
			value.Elem(),
		)
	}
	var dstValue reflect.Value
	var ok bool
	if dstValue, ok = dst.(reflect.Value); !ok {
		dstValue = reflect.ValueOf(dst)
	}
	dstValue.Elem().Set(sliceValue)
	return nil
}

func (m *query) Where(args ...interface{}) db.IQuery {
	if len(args) == 0 {
		return m
	}

	if f, ok := args[0].(bson.M); ok {
		m.filter = f
	}
	return m
}

func (m *query) sort(flag int, fields []string) {
	underscore.Chain(fields).Each(func(r string, _ int) {
		m.sorts = append(m.sorts, bson.E{
			Key:   r,
			Value: flag,
		})
	})
}

func newQuery(pool *connectPool, modelStruct identity.IStruct) *query {
	q := &query{
		modelStruct: modelStruct,
		pool:        pool,
	}
	q.Reset()
	return q
}
