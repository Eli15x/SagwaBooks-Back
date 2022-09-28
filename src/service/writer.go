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
	CreateNewWriter(ctx context.Context, email string, password string, birthDate string, name string, city string, rg string, cpf string, telefone string, image string) (string, error)
	EditWriter(ctx context.Context, writerId string, email string, password string, birthDate string, name string, city string, rg string, cpf string, telefone string, image string) error
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

func (w *writer) CreateNewWriter(ctx context.Context, email string, password string, birthDate string, name string, city string, rg string, cpf string, telefone string, image string) (string, error) {

	//na hora de cadastrar conferir se já existe usuario com aquele email.
	var writerId = utils.CreateCodeId()
	writer := &models.Writer{
		WriterId:  writerId,
		Name:      name,
		Email:     email,
		PassWord:  password,
		Telefone:  telefone,
		Rg:        rg,
		Cpf:       cpf,
		City:      city,
		BirthDate: birthDate,
		Image:     image,
	}

	writerInsert := structs.Map(writer)

	_, err := client.GetInstance().Insert(ctx, "writer", writerInsert)
	if err != nil {
		return "", errors.New("Create New Writer: problem to insert into MongoDB")
	}

	return writerId, nil
}

/*func (w *writer) AddAdress(ctx context.Context, userId string, rua string, complemento string, numero string, bairro string, cidade string, cep string) error {

	user := map[string]interface{}{
		"UserId":      userId,
		"Rua":         rua,
		"Complemento": complemento,
		"Numero":      numero,
		"Bairro":      bairro,
		"Cidade":      cidade,
		"Cep":         cep,
	}

	userUpdate := structs.Map(user)
	change := bson.M{"$set": userUpdate}

	UserId := map[string]interface{}{"UserId": userId}
	_, err := client.GetInstance().UpdateOne(ctx, "user", UserId, change)
	if err != nil {
		return errors.New("add Adress User: problem to uptade into MongoDB")
	}

	return nil
}*/

func (w *writer) EditWriter(ctx context.Context, writerId string, email string, password string, birthDate string, name string, city string, rg string, cpf string, telefone string, image string) error {

	//ver logica para o bookId porque a pessoa conseguiria alterar o bookId nesse caso...
	writer := &models.Writer{ //mudar depois para interface normal.
		WriterId:  writerId,
		Name:      name,
		Email:     email,
		PassWord:  password,
		Telefone:  telefone,
		Rg:        rg,
		Cpf:       cpf,
		City:      city,
		BirthDate: birthDate,
		Image:     image,
	}

	writerUpdate := structs.Map(writer)
	change := bson.M{"$set": writerUpdate}

	write := map[string]interface{}{"WriterId": writerId}
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

	result, err := repository.Find(ctx, "user", WriterId, &Writer)
	if err != nil {
		return nil, errors.New("Get Writer error: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (w *writer) GetInformationWriters(ctx context.Context) ([]bson.M, error) {
	var Writer models.Writer

	WriterId := map[string]interface{}{"WriterId": ""}

	result, err := repository.Find(ctx, "writer", WriterId, &Writer)
	if err != nil {
		return nil, errors.New("Get Writer error: problem to Find Id into MongoDB")
	}

	return result, nil
}
