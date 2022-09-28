package service

import (
	"context"

	testifyMock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type UserMock struct {
	testifyMock.Mock
}

func (u *UserMock) CreateNewUser(ctx context.Context, name string, email string, password string, telefone string) (string, error) {
	args := u.Called(ctx, name, email, password, telefone)
	return args.Get(0).(string), args.Error(1)
}

func (u *UserMock) EditUser(ctx context.Context, userId string, name string, email string, password string, telefone string) error {
	args := u.Called(ctx, userId, name, email, password, telefone)
	return args.Error(0)
}

func (u *UserMock) DeleteUser(ctx context.Context, userId string) error {
	args := u.Called(ctx, userId)
	return args.Error(0)
}

func (u *UserMock) GetInformationUser(ctx context.Context, id string) ([]bson.M, error) {
	args := u.Called(ctx, id)
	return args.Get(0).([]bson.M), args.Error(1)
}

func (u *UserMock) ValidateUser(ctx context.Context, email string, password string) ([]bson.M, error) {
	args := u.Called(ctx, email, password)
	return args.Get(0).([]bson.M), args.Error(1)
}

func (u *UserMock) AddAdress(ctx context.Context, userId string, rua string, complemento string, numero string, bairro string, cidade string, cep string) error {
	args := u.Called(ctx, userId, rua, complemento, numero, bairro, cidade, cep)
	return args.Error(0)
}
