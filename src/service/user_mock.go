package service

import (
	"context"

	testifyMock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type CardMock struct {
	testifyMock.Mock
}

func (c *CardMock) CreateNewUser(ctx context.Context, name string, email string, password string, telefone string) (string, error) {
	args := c.Called(ctx, name, email, password, telefone)
	return args.Get(0).(string), args.Error(1)
}

func (c *CardMock) EditUser(ctx context.Context, userId string, name string, email string, password string, telefone string) error {
	args := c.Called(ctx, userId, name, email, password, telefone)
	return args.Error(0)
}

func (c *CardMock) DeleteUser(ctx context.Context, userId string) error {
	args := c.Called(ctx, userId)
	return args.Error(0)
}

func (c *CardMock) GetInformationUser(ctx context.Context, id string) ([]bson.M, error) {
	args := c.Called(ctx, id)
	return args.Get(0).([]bson.M), args.Error(1)
}

func (c *CardMock) ValidateUser(ctx context.Context, email string, password string) ([]bson.M, error) {
	args := c.Called(ctx, email, password)
	return args.Get(0).([]bson.M), args.Error(1)
}

func (c *CardMock) AddAdress(ctx context.Context, userId string, rua string, complemento string, numero string, bairro string, cidade string, cep string) error {
	args := c.Called(ctx, userId, rua, complemento, numero, bairro, cidade, cep)
	return args.Error(0)
}
