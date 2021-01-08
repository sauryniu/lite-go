package mongodb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type query struct {
	Filter bson.M
	Opt    *options.FindOptions
	Pool   *connectPool
	Sort   bson.D
	Struct identity.IStruct
}

func (m query) Count() (int64, error) {
	c, err := m.Pool.GetCollection(m.Struct)
	if err != nil {
		return 0, err
	}

	defer m.Reset()
	return c.CountDocuments(m.Pool.Ctx, m.Filter)
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
	m.Filter = make(bson.M)
	m.Opt = options.Find()
	m.Sort = bson.D{}
}

func (m *query) Skip(v int) db.IQuery {
	m.Opt = m.Opt.SetSkip(
		int64(v),
	)
	return m
}

func (m *query) Take(v int) db.IQuery {
	m.Opt = m.Opt.SetLimit(
		int64(v),
	)
	return m
}

func (m *query) ToArray(dst interface{}) error {
	c, err := m.Pool.GetCollection(m.Struct)
	if err != nil {
		return err
	}

	defer m.Reset()

	if len(m.Sort) > 0 {
		m.Opt = m.Opt.SetSort(m.Sort)
	}

	cur, err := c.Find(m.Pool.Ctx, m.Filter, m.Opt)
	if err != nil {
		return err
	}

	sliceType := reflect.SliceOf(
		m.Struct.GetType(),
	)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	for cur.Next(m.Pool.Ctx) {
		value := reflect.New(
			m.Struct.GetType(),
		)
		temp := value.Interface()
		err = cur.Decode(temp)

		sliceValue = reflect.Append(
			sliceValue,
			value.Elem(),
		)
	}
	reflect.ValueOf(dst).Elem().Set(sliceValue)
	return nil
}

func (m *query) Where(args ...interface{}) db.IQuery {
	if len(args) == 0 {
		return m
	}

	if f, ok := args[0].(bson.M); ok {
		m.Filter = f
	}
	return m
}

func (m *query) sort(flag int, fields []string) {
	underscore.Chain(fields).Each(func(r string, _ int) {
		m.Sort = append(m.Sort, bson.E{
			Key:   r,
			Value: flag,
		})
	})
}

func newQuery(pool *connectPool, s identity.IStruct) *query {
	q := &query{
		Pool:   pool,
		Struct: s,
	}
	q.Reset()
	return q
}
