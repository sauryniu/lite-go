package mongosvc

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type unitOfWork struct {
	addQueue    []contract.IDbModel
	removeQueue []contract.IDbModel
	saveQueue   []contract.IDbModel
	client      *clientWrapper
}

func (m *unitOfWork) Commit() error {
	if len(m.addQueue) == 0 && len(m.removeQueue) == 0 && len(m.saveQueue) == 0 {
		return nil
	}

	client, err := m.client.GetClient()
	if err != nil {
		return err
	}

	return client.UseSession(m.client.Ctx, func(ctx mongo.SessionContext) (err error) {
		if err = m.commitAdd(ctx); err != nil {
			return err
		}

		if err = m.commitRemove(ctx); err != nil {
			return err
		}

		return m.commitSave(ctx)
	})
}

func (m *unitOfWork) RegisterAdd(entry contract.IDbModel) {
	m.addQueue = append(m.addQueue, entry)
}

func (m *unitOfWork) RegisterRemove(entry contract.IDbModel) {
	m.removeQueue = append(m.removeQueue, entry)
}

func (m *unitOfWork) RegisterSave(entry contract.IDbModel) {
	m.saveQueue = append(m.saveQueue, entry)
}

func (m *unitOfWork) commitAdd(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.addQueue).Map(func(r contract.IDbModel, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		model := getModelMetadata(
			entryValue.Type(),
		)
		col, rErr := m.client.GetCollection(model)
		if rErr != nil {
			return rErr
		}

		doc := make(bson.M)
		underscore.Chain(
			model.FindFields(),
		).Each(func(r *fieldMetadata, _ int) {
			doc[r.GetColumnName()] = r.GetValue(entryValue)
		})
		_, rErr = col.InsertOne(ctx, doc)
		return rErr
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	m.addQueue = make([]contract.IDbModel, 0)
	return
}

func (m *unitOfWork) commitRemove(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.removeQueue).Map(func(r contract.IDbModel, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		model := getModelMetadata(
			entryValue.Type(),
		)
		col, rErr := m.client.GetCollection(model)
		if rErr != nil {
			return rErr
		}

		idField, rErr := model.GetIDField()
		if rErr != nil {
			return rErr
		}

		filter := make(bson.M)
		filter[idField.GetColumnName()] = idField.GetValue(entryValue)
		_, rErr = col.DeleteOne(ctx, filter)
		return rErr
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	m.removeQueue = make([]contract.IDbModel, 0)
	return
}

func (m *unitOfWork) commitSave(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.saveQueue).Map(func(r contract.IDbModel, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		model := getModelMetadata(
			entryValue.Type(),
		)
		col, rErr := m.client.GetCollection(model)
		if rErr != nil {
			return rErr
		}

		filer := make(bson.M)
		doc := make(bson.M)
		underscore.Chain(
			model.FindFields(),
		).Each(func(r *fieldMetadata, _ int) {
			if r.GetTableName() != "" {
				filer[r.GetColumnName()] = r.GetValue(entryValue)
			}
			doc[r.GetColumnName()] = r.GetValue(entryValue)
		})
		_, rErr = col.UpdateOne(ctx, filer, bson.M{
			"$set": doc,
		})
		return rErr
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	m.saveQueue = make([]contract.IDbModel, 0)
	return
}

func newUnitOfWork(client *clientWrapper) *unitOfWork {
	return &unitOfWork{
		addQueue:    make([]contract.IDbModel, 0),
		removeQueue: make([]contract.IDbModel, 0),
		saveQueue:   make([]contract.IDbModel, 0),
		client:      client,
	}
}
