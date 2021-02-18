package mongodb

import (
	"context"

	"github.com/ahl5esoft/lite-go/db/identity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connectPool struct {
	client      *mongo.Client
	ctx         context.Context
	collections map[string]*mongo.Collection
	db          *mongo.Database
	dbName      string
	option      *options.ClientOptions
}

func (m *connectPool) GetClient() (*mongo.Client, error) {
	if m.client == nil {
		var err error
		m.client, err = mongo.Connect(m.ctx, m.option)
		if err != nil {
			return nil, err
		}
	}
	return m.client, nil
}

func (m *connectPool) GetCollection(s identity.IStruct) (*mongo.Collection, error) {
	name, err := s.GetName()
	if err != nil {
		return nil, err
	}

	if _, ok := m.collections[name]; !ok {
		db, err := m.GetDb()
		if err != nil {
			return nil, err
		}

		m.collections[name] = db.Collection(name)
	}

	return m.collections[name], nil
}

func (m *connectPool) GetDb() (*mongo.Database, error) {
	if m.db == nil {
		client, err := m.GetClient()
		if err != nil {
			return nil, err
		}

		m.db = client.Database(m.dbName)
	}

	return m.db, nil
}

func newPool(uri, name string) *connectPool {
	return &connectPool{
		collections: make(map[string]*mongo.Collection),
		ctx:         context.Background(),
		dbName:      name,
		option:      options.Client().ApplyURI(uri),
	}
}
