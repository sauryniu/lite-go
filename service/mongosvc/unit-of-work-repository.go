package mongosvc

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/contract"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type unitOfWorkRepository struct {
	driverFactory *driverFactory
	deleteQueue   []contract.IDbModel
	insertQueue   []contract.IDbModel
	updateQueue   []contract.IDbModel
}

func (m *unitOfWorkRepository) Commit() error {
	if len(m.insertQueue) == 0 && len(m.deleteQueue) == 0 && len(m.updateQueue) == 0 {
		return nil
	}

	client, err := m.driverFactory.BuildClient()
	if err != nil {
		return err
	}

	return client.UseSession(m.driverFactory.Ctx, func(ctx mongo.SessionContext) (err error) {
		if err = m.commitInsert(ctx); err != nil {
			return err
		}

		if err = m.commitDelete(ctx); err != nil {
			return err
		}

		return m.commitUpdate(ctx)
	})
}

func (m *unitOfWorkRepository) RegisterDelete(entry contract.IDbModel) {
	m.deleteQueue = append(m.deleteQueue, entry)
}

func (m *unitOfWorkRepository) RegisterInsert(entry contract.IDbModel) {
	m.insertQueue = append(m.insertQueue, entry)
}

func (m *unitOfWorkRepository) RegisterUpdate(entry contract.IDbModel) {
	m.updateQueue = append(m.updateQueue, entry)
}

func (m *unitOfWorkRepository) commitDelete(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.deleteQueue).Map(func(r contract.IDbModel, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		model := getModelMetadata(
			entryValue.Type(),
		)
		col, rErr := m.driverFactory.BuildCollection(model)
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
	m.deleteQueue = make([]contract.IDbModel, 0)
	return
}

func (m *unitOfWorkRepository) commitInsert(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.insertQueue).Map(func(r contract.IDbModel, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		model := getModelMetadata(
			entryValue.Type(),
		)
		col, rErr := m.driverFactory.BuildCollection(model)
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
	m.insertQueue = make([]contract.IDbModel, 0)
	return
}

func (m *unitOfWorkRepository) commitUpdate(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.updateQueue).Map(func(r contract.IDbModel, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		model := getModelMetadata(
			entryValue.Type(),
		)
		col, rErr := m.driverFactory.BuildCollection(model)
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
	m.updateQueue = make([]contract.IDbModel, 0)
	return
}

func newUnitOfWorkRepository(driverFactory *driverFactory) *unitOfWorkRepository {
	return &unitOfWorkRepository{
		driverFactory: driverFactory,
		deleteQueue:   make([]contract.IDbModel, 0),
		insertQueue:   make([]contract.IDbModel, 0),
		updateQueue:   make([]contract.IDbModel, 0),
	}
}
