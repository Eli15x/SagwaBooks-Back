package client

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/labstack/gommon/log"
)

var (
	once          sync.Once
	mongoInstance MongoDB
)

type MongoDB interface {
	Insert(ctx context.Context, collName string, doc interface{}) (interface{}, error)
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.Cursor, error)
	FindAll(ctx context.Context, collName string, doc interface{}) (*mongo.Cursor, error)
	FindOne(ctx context.Context, collName string, filter interface{}, opts ...*options.FindOneOptions) MongoFindOneResult
	Count(ctx context.Context, collName string, query map[string]interface{}) (int64, error)
	UpdateOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.UpdateResult, error)
	Remove(ctx context.Context, collName string, query map[string]interface{}) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
	Initialize(ctx context.Context) error
	Disconnect()
}

type MongoFindOneResult interface {
	Err() error
	Decode(v interface{}) error
}

type mongodbImpl struct {
	client *mongo.Client
	dbName string
}

func GetInstance() MongoDB {
	once.Do(func() {
		mongoInstance = &mongodbImpl{}
	})
	return mongoInstance
}

func (m *mongodbImpl) Initialize(ctx context.Context) error {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://elisacds:elisacds@cluster0.e7uxp.mongodb.net/SagwaBooks?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	m.dbName = "SagwaBooks"
	m.client = client
	return nil
}

func (m *mongodbImpl) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return m.client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		err = fn(sessionContext)
		if err != nil {
			return sessionContext.AbortTransaction(sessionContext)
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
}

// Insert stores documents in the collection
func (m *mongodbImpl) Insert(ctx context.Context, collName string, doc interface{}) (interface{}, error) {

	insertedObject, err := m.client.Database(m.dbName).Collection(collName).InsertOne(ctx, doc)
	if insertedObject == nil {
		fmt.Println(err)
		return nil, err
	}
	return insertedObject.InsertedID, err
}

// Find finds all documents in the collection
func (m *mongodbImpl) Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.Cursor, error) {
	cur, err := m.client.Database(m.dbName).Collection(collName).Find(ctx, query)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func (m *mongodbImpl) FindAll(ctx context.Context, collName string, doc interface{}) (*mongo.Cursor, error) {
	cur, err := m.client.Database(m.dbName).Collection(collName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	return cur, nil
}

// FindOne finds one document in mongo
func (m *mongodbImpl) FindOne(ctx context.Context, collName string, filter interface{}, opts ...*options.FindOneOptions) MongoFindOneResult {
	coll := m.client.Database(m.dbName).Collection(collName)
	findResult := coll.FindOne(ctx, filter)
	return findResult
}

// UpdateOne updates one or more documents in the collection
func (m *mongodbImpl) UpdateOne(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error) {

	updateResult, err := m.client.Database(m.dbName).Collection(collName).UpdateOne(ctx, selector, update)
	fmt.Println(err)
	return updateResult, err
}

// Remove one or more documents in the collection
func (m *mongodbImpl) Remove(ctx context.Context, collName string, selector map[string]interface{}) error {
	_, err := m.client.Database(m.dbName).Collection(collName).DeleteOne(ctx, selector)
	return err
}

// Count returns the number of documents of the query
func (m *mongodbImpl) Count(ctx context.Context, collName string, query map[string]interface{}) (int64, error) {
	return m.client.Database(m.dbName).Collection(collName).CountDocuments(ctx, query)
}

func (m *mongodbImpl) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_ = m.client.Disconnect(ctx)
}
