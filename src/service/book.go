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
	instanceBook CommandBook
	onceBook     sync.Once
)

type CommandBook interface {
	CreateNewBook(ctx context.Context, name string, autor string, genero string, preco string, prioridade string) error
	EditBook(ctx context.Context, bookId string, name string, autor string, genero string, preco string, prioridade string) error
	DeleteBook(ctx context.Context, bookId string) error
	GetInformationByName(ctx context.Context, name string) ([]bson.M, error)
	GetInformationByAutor(ctx context.Context, autor string) ([]bson.M, error)
	GetInformationByGenero(ctx context.Context, genero string) ([]bson.M, error)
	GetInformationByPriority(ctx context.Context, priority string) ([]bson.M, error)
}

type book struct{}

func GetInstanceBook() CommandBook {
	onceBook.Do(func() {
		instanceBook = &book{}
	})
	return instanceBook
}

func (b *book) CreateNewBook(ctx context.Context, name string, autor string, genero string, preco string, prioridade string) error {

	//na hora de cadastrar conferir se já existe usuario com aquele email.
	var bookId = utils.CreateCodeId()
	book := &models.Book{
		BookId:     bookId,
		Name:       name,
		Autor:      autor,
		Genero:     genero,
		Preco:      preco,
		Prioridade: prioridade,
	}

	bookInsert := structs.Map(book)

	_, err := client.GetInstance().Insert(ctx, "book", bookInsert)
	if err != nil {
		return errors.New("Create New book: problem to insert into MongoDB")
	}

	return nil
}

func (b *book) EditBook(ctx context.Context, bookId string, name string, autor string, genero string, preco string, prioridade string) error {

	//ver logica para o bookId porque a pessoa conseguiria alterar o bookId nesse caso...
	book := &models.Book{ //mudar depois para interface normal.
		BookId:     bookId,
		Name:       name,
		Autor:      autor,
		Genero:     genero,
		Preco:      preco,
		Prioridade: prioridade,
	}

	bookUpdate := structs.Map(book)
	change := bson.M{"$set": bookUpdate}

	BookId := map[string]interface{}{"BookId": bookId}
	_, err := client.GetInstance().UpdateOne(ctx, "book", BookId, change)
	if err != nil {
		return errors.New("Edit Book: problem to uptade into MongoDB")
	}

	return nil
}

func (b *book) DeleteBook(ctx context.Context, bookId string) error {

	BookId := map[string]interface{}{"BookId": bookId}
	err := client.GetInstance().Remove(ctx, "book", BookId)
	if err != nil {
		return errors.New("Edit Book: problem to uptade into MongoDB")
	}

	return nil
}

func (b *book) GetInformationByName(ctx context.Context, name string) ([]bson.M, error) {
	var book models.Book

	Name := map[string]interface{}{"Name": name}

	result, err := repository.Find(ctx, "book", Name, &book)
	if err != nil {
		return nil, errors.New("Get Information By Name Book: problem to Find name into MongoDB")
	}

	return result, nil
}

func (b *book) GetInformationByAutor(ctx context.Context, autor string) ([]bson.M, error) {
	var book models.Book

	Autor := map[string]interface{}{"Autor": autor}

	result, err := repository.Find(ctx, "book", Autor, &book)
	if err != nil {
		return nil, errors.New("Get Information By Autor Book: problem to Find autor into MongoDB")
	}

	return result, nil
}

func (b *book) GetInformationByGenero(ctx context.Context, genero string) ([]bson.M, error) {
	var book models.Book

	Genero := map[string]interface{}{"Genero": genero}

	result, err := repository.Find(ctx, "book", Genero, &book)
	if err != nil {
		return nil, errors.New("Get Information By Genero Book: problem to Find genero into MongoDB")
	}

	return result, nil
}

func (b *book) GetInformationByPriority(ctx context.Context, priority string) ([]bson.M, error) {
	var book models.Book

	Prioridade := map[string]interface{}{"Prioridade": priority}

	result, err := repository.Find(ctx, "book", Prioridade, &book)
	if err != nil {
		return nil, errors.New("Get Information By Priority Book: problem to Find priority level into MongoDB")
	}

	return result, nil
}
