package service

import (
	"context"
	"errors"
	"sync"

	"github.com/Eli15x/SagwaBooks-Back/src/client"
	"github.com/Eli15x/SagwaBooks-Back/src/models"
	"github.com/Eli15x/SagwaBooks-Back/src/repository"
	"github.com/Eli15x/SagwaBooks-Back/utils"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceWriter CommandWriter
	onceWriter     sync.Once
)

type CommandWriter interface {
	CreateNewWriter(ctx context.Context, writer *models.Writer) (string, error)
	EditWriter(ctx context.Context, writer *models.Writer) error
	GetInformationWriter(ctx context.Context, id string) ([]bson.M, error)
	DeleteWriter(ctx context.Context, writerId string) error
	GetInformationWriters(ctx context.Context) ([]bson.M, error)
	//AddAdress(ctx context.Context, userId string, rua string, complemento string, numero string, bairro string, cidade string, cep string) error
}

type writer struct{}

func GetInstanceWriter() CommandWriter {
	onceWriter.Do(func() {
		instanceWriter = &writer{}
	})
	return instanceWriter
}

func (w *writer) CreateNewWriter(ctx context.Context, writer *models.Writer) (string, error) {

	var writerId = utils.CreateCodeId()
	writer.WriterId = writerId

	writerInsert := structs.Map(writer)

	_, err := client.GetInstance().Insert(ctx, "writer", writerInsert)
	if err != nil {
		return "", errors.New("Create New Writer: problem to insert into MongoDB")
	}

	return writerId, nil
}

func (w *writer) EditWriter(ctx context.Context, writer *models.Writer) error {

	writerUpdate := structs.Map(writer)
	change := bson.M{"$set": writerUpdate}

	write := map[string]interface{}{"WriterId": writer.WriterId}
	_, err := client.GetInstance().UpdateOne(ctx, "writer", write, change)
	if err != nil {
		return errors.New("Edit Write: problem to uptade into MongoDB")
	}

	return nil
}

func (w *writer) DeleteWriter(ctx context.Context, writerId string) error {

	WriterId := map[string]interface{}{"WriterId": writerId}
	err := client.GetInstance().Remove(ctx, "writer", WriterId)
	if err != nil {
		return errors.New("Delete Writer Error: problem to Delete into MongoDB")
	}

	return nil
}

func (w *writer) GetInformationWriter(ctx context.Context, id string) ([]bson.M, error) {
	var Writer models.Writer

	WriterId := map[string]interface{}{"WriterId": id}

	result, err := repository.Find(ctx, "writer", WriterId, &Writer)
	if err != nil {
		return nil, errors.New("Get Writer error: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (w *writer) GetInformationWriters(ctx context.Context) ([]bson.M, error) {
	var Writer models.Writer

	result, err := repository.FindAll(ctx, "writer", &Writer)
	if err != nil {
		return nil, errors.New("Get Writer error: problem to Find Id into MongoDB")
	}

	return result, nil
}
