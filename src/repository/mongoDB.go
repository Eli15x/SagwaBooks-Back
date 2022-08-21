package repository

import (
	"context"
	"errors"

	"github.com/Eli15x/SagwaBooks-Back/src/client"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoDB interface {
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error)
}

func Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query, doc)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDb")
	}

	var content []bson.M
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}
