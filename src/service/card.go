package service

import (
	"context"
	"errors"
	"sync"

	"github.com/Eli15x/SagwaBooks-Back/src/client"
	"github.com/Eli15x/SagwaBooks-Back/src/models"
	"github.com/Eli15x/SagwaBooks-Back/src/repository"
	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceCard CommandCard
	onceCard     sync.Once
)

type CommandCard interface {
	CreateNewCard(ctx context.Context, userId string, numero string, name string, mounth string, year string) error
	EditCard(ctx context.Context, userId string, numero string, name string, mounth string, year string) error
	DeleteCard(ctx context.Context, numero string) error
	GetInformationByUserId(ctx context.Context, userId string) ([]bson.M, error)
}

type Card struct{}

func GetInstanceCard() CommandCard {
	onceCard.Do(func() {
		instanceCard = &Card{}
	})
	return instanceCard
}

func (c *Card) CreateNewCard(ctx context.Context, userId string, numero string, name string, mounth string, year string) error {
	Card := &models.Card{
		UserId: userId,
		Numero: numero,
		Name:   name,
		Mounth: mounth,
		Year:   year,
	}

	CardInsert := structs.Map(Card)

	_, err := client.GetInstance().Insert(ctx, "card", CardInsert)
	if err != nil {
		return errors.New("Create New Card: problem to insert into MongoDB")
	}

	return nil
}

func (c *Card) EditCard(ctx context.Context, userId string, numero string, name string, mounth string, year string) error {

	//ver logica para o CardId porque a pessoa conseguiria alterar o CardId nesse caso...
	Card := &models.Card{ //mudar depois para interface normal.
		UserId: userId,
		Numero: numero,
		Name:   name,
		Mounth: mounth,
		Year:   year,
	}

	CardUpdate := structs.Map(Card)
	change := bson.M{"$set": CardUpdate}

	Numero := map[string]interface{}{"Numero": numero}
	_, err := client.GetInstance().UpdateOne(ctx, "Card", Numero, change)
	if err != nil {
		return errors.New("Edit Card: problem to uptade into MongoDB")
	}

	return nil
}

func (c *Card) DeleteCard(ctx context.Context, numero string) error {

	Numero := map[string]interface{}{"Numero": numero}
	err := client.GetInstance().Remove(ctx, "card", Numero)
	if err != nil {
		return errors.New("Delete Card: problem to uptade into MongoDB")
	}

	return nil
}

func (c *Card) GetInformationByUserId(ctx context.Context, userId string) ([]bson.M, error) {
	var Card models.Card

	UserId := map[string]interface{}{"UserId": userId}

	result, err := repository.Find(ctx, "Card", UserId, &Card)
	if err != nil {
		return nil, errors.New("Get Information By userId Card: problem to Find name into MongoDB")
	}

	return result, nil
}
