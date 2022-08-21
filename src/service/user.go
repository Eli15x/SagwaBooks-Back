package service

import (
	"context"
	"errors"
	"sync"

	storage "github.com/Eli15x/SagwaBooks-Back/src/client"
	"github.com/Eli15x/SagwaBooks-Back/src/models"
	"github.com/Eli15x/SagwaBooks-Back/src/repository"
	"github.com/Eli15x/SagwaBooks-Back/utils"
	"github.com/fatih/structs"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceUser CommandUser
	onceUser     sync.Once
)

type CommandUser interface {
	CreateNewUser(ctx context.Context, name string, email string, password string, telefone string) (string, error)
	EditUser(ctx context.Context, userId string, name string, email string, password string, telefone string) error
	DeleteUser(ctx context.Context, userId string) error
	GetInformationUser(ctx context.Context, id string) ([]bson.M, error)
	ValidateUser(ctx context.Context, email string, password string) ([]bson.M, error)
}

type user struct{}

func GetInstanceUser() CommandUser {
	onceUser.Do(func() {
		instanceUser = &user{}
	})
	return instanceUser
}

func (u *user) CreateNewUser(ctx context.Context, name string, email string, password string, telefone string) (string, error) {

	//na hora de cadastrar conferir se já existe usuario com aquele email.
	var userId = utils.CreateCodeId()
	user := &models.User{
		UserId:   userId,
		Name:     name,
		Email:    email,
		PassWord: password,
		Telefone: telefone,
	}

	userInsert := structs.Map(user)

	_, err := storage.GetInstance().Insert(ctx, "user", userInsert)
	if err != nil {
		return "", errors.New("Create New User: problem to insert into MongoDB")
	}

	return userId, nil
}

func (u *user) EditUser(ctx context.Context, userId string, name string, email string, password string, telefone string) error {

	//ver logica para o bookId porque a pessoa conseguiria alterar o bookId nesse caso...
	user := &models.User{ //mudar depois para interface normal.
		UserId:   userId,
		Email:    email,
		Name:     name,
		PassWord: password,
		Telefone: telefone,
	}

	userUpdate := structs.Map(user)
	change := bson.M{"$set": userUpdate}

	UserId := map[string]interface{}{"UserId": userId}
	_, err := storage.GetInstance().UpdateOne(ctx, "user", UserId, change)
	if err != nil {
		return errors.New("Edit User: problem to uptade into MongoDB")
	}

	return nil
}

func (u *user) DeleteUser(ctx context.Context, userId string) error {

	UserId := map[string]interface{}{"UserId": userId}
	err := storage.GetInstance().Remove(ctx, "user", UserId)
	if err != nil {
		return errors.New("Edit User: problem to uptade into MongoDB")
	}

	return nil
}

func (u *user) GetInformationUser(ctx context.Context, id string) ([]bson.M, error) {
	var user models.User

	userId := map[string]interface{}{"UserId": id}

	result, err := repository.Find(ctx, "user", userId, &user)
	if err != nil {
		return nil, errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (u *user) ValidateUser(ctx context.Context, email string, password string) ([]bson.M, error) {

	var user *models.User

	filter := map[string]interface{}{"Email": email, "PassWord": password}
	result, err := repository.Find(ctx, "user", filter, &user)
	if err != nil {
		return nil, errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	log.Infof("[CheckInformationValid] Object : %s \n", user, "")

	return result, nil
}
