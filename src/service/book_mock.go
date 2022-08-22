package service

import (
	"context"

	testifyMock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type BookMock struct {
	testifyMock.Mock
}

func (b *BookMock) CreateNewBook(ctx context.Context, name string, autor string, genero string, preco string, prioridade string) error {
	args := b.Called(ctx, name, autor, genero, preco, prioridade)
	return args.Error(0)
}

func (b *BookMock) EditBook(ctx context.Context, bookId string, name string, autor string, genero string, preco string, prioridade string) error {
	args := b.Called(ctx, bookId, name, autor, genero, preco, prioridade)
	return args.Error(0)
}

func (b *BookMock) DeleteBook(ctx context.Context, bookId string) error {
	args := b.Called(ctx, bookId)
	return args.Error(0)
}

func (b *BookMock) GetInformationByName(ctx context.Context, name string) ([]bson.M, error) {
	args := b.Called(ctx, name)
	return args.Get(0).([]bson.M), args.Error(1)
}
