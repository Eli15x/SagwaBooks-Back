package client

import (
	"context"

	testifyMock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBMock struct {
	testifyMock.Mock
}

func (m *MongoDBMock) Initialize(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MongoDBMock) Insert(ctx context.Context, collName string, doc interface{}) (interface{}, error) {
	args := m.Called(ctx, collName, doc)
	return args.Get(0).(interface{}), args.Error(1)
}

func (m *MongoDBMock) Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.Cursor, error) {
	args := m.Called(ctx, collName, query, doc)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MongoDBMock) FindOne(ctx context.Context, collName string, filter interface{}, opts ...*options.FindOneOptions) MongoFindOneResult {
	args := m.Called(ctx, collName, filter)
	return args.Get(0).(MongoFindOneResult)
}

func (m *MongoDBMock) Count(ctx context.Context, collName string, query map[string]interface{}) (int64, error) {
	args := m.Called(ctx, collName, query)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MongoDBMock) UpdateOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, collName, query, doc)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MongoDBMock) Remove(ctx context.Context, collName string, query map[string]interface{}) error {
	args := m.Called(ctx, collName, query)
	return args.Error(0)
}

func (m *MongoDBMock) Disconnect() {
	return
}
