package mongodb

import (
	"context"

	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connectPool struct {
	Ctx         context.Context
	Name        string
	Collections map[string]*mongo.Collection
	URI         string

	client *mongo.Client
	db     *mongo.Database
}

func (m *connectPool) GetClient() (*mongo.Client, error) {
	if m.client == nil {
		opt := options.Client().ApplyURI(m.URI)
		var err error
		m.client, err = mongo.Connect(m.Ctx, opt)
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

	if _, ok := m.Collections[name]; !ok {
		db, err := m.GetDb()
		if err != nil {
			return nil, err
		}

		m.Collections[name] = db.Collection(name)
	}

	return m.Collections[name], nil
}

func (m *connectPool) GetDb() (*mongo.Database, error) {
	if m.db == nil {
		client, err := m.GetClient()
		if err != nil {
			return nil, err
		}

		m.db = client.Database(m.Name)
	}

	return m.db, nil
}

func newPool(name string, uri string) *connectPool {
	return &connectPool{
		Collections: make(map[string]*mongo.Collection),
		Ctx:         context.Background(),
		Name:        name,
		URI:         uri,
	}
}
