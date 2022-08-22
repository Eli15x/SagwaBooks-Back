package service

import (
	"context"

	testifyMock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type CardMock struct {
	testifyMock.Mock
}

func (c *CardMock) CreateNewCard(ctx context.Context, userId string, numero string, name string, mounth string, year string) error {
	args := c.Called(ctx, userId, numero, name, mounth, year)
	return args.Error(0)
}

func (c *CardMock) EditCard(ctx context.Context, userId string, numero string, name string, mounth string, year string) error {
	args := c.Called(ctx, userId, numero, name, mounth, year)
	return args.Error(0)
}

func (c *CardMock) DeleteCard(ctx context.Context, numero string) error {
	args := c.Called(ctx, numero)
	return args.Error(0)
}

func (c *CardMock) GetInformationByUserId(ctx context.Context, userId string) ([]bson.M, error) {
	args := c.Called(ctx, userId)
	return args.Get(0).([]bson.M), args.Error(1)
}
