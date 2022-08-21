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
	//AddInformationProfile(ctx context.Context, id string, job []string, message string) error
	GetInformationUser(ctx context.Context, id string) ([]bson.M, error)
	ValidateUser(ctx context.Context, email string, password string) ([]bson.M, error)
	/*AddRelationFriendProfile(ctx context.Context, UserId_user string, UserId_value string, friend *models.Friend) error
	AddRequestFriend(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error
	DeleteFriendRequest(ctx context.Context, UserId string, FriendId string, friendUser *models.Friend) error
	AddContent(ctx context.Context, id string, content string) error*/
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

/*func (u *user) InformationUserCard(ctx context.Context, name string, number string, data string,) error {
	userId := map[string]interface{}{"UserId": id}

	//existe com aquele id
	mgoErr := storage.GetInstance().FindOne(ctx, "profile", userId)
	if mgoErr != nil {
		return errors.New("Add Information Profile: problem to Find Id into MongoDB")
	}

	profileUpdate := map[string]interface{}{
		"Job":            job,
		"ProfileMessage": message,
	}

	fmt.Println(profileUpdate)

	change := bson.M{"$set": profileUpdate}

	_, err := storage.GetInstance().UpdateOne(ctx, "profile", userId, change)
	if err != nil {
		return errors.New("Create New Profile: problem to update into MongoDB")
	}

	return nil
}*/

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
