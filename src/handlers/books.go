package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/SagwaBooks-Back/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateBook(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	autor := json_map["autor"].(string)
	genero := json_map["genero"].(string)
	prioridade := json_map["prioridade"].(string)
	preco := json_map["preco"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	if autor == "" {
		c.String(400, "Create User Error: email not find")
		return
	}

	if genero == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	if prioridade == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	if preco == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	err = service.GetInstanceBook().CreateNewBook(context.Background(), name, autor, genero, preco, prioridade)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func EditBook(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	bookId := json_map["bookId"].(string)
	name := json_map["name"].(string)
	autor := json_map["autor"].(string)
	genero := json_map["genero"].(string)
	prioridade := json_map["prioridade"].(string)
	preco := json_map["preco"].(string)

	if bookId == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	if name == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	if autor == "" {
		c.String(400, "Create User Error: email not find")
		return
	}

	if genero == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	if prioridade == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	if preco == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	err = service.GetInstanceBook().EditBook(context.Background(), bookId, name, autor, genero, preco, prioridade)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteBook(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	bookId := json_map["bookId"].(string)

	if bookId == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	err = service.GetInstanceBook().DeleteBook(context.Background(), bookId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

/*
func EditUserPermission(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	permission := json_map["permission"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.
	userId := json_map["userId"].(string)         //está dando erro quando tenta pegar o "email" e ele não existe.

	if userId == "" {
		c.String(400, "Edit User Permission Error: userId not find")
		return
	}

	if permission == "" {
		c.String(400, "Edit User Permission Error: permission not find")
		return
	}

	//ver se permission é um numero valido (0,1,2,3)

	err = service.GetInstanceUser().EditPermissionUser(context.Background(), userId, permission)
	if err != nil {
		c.String(403, err.Error())
		return
	}

	c.String(http.StatusOK, "Ok")
}*/

func GetBookByName(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)

	result, err := service.GetInstanceBook().GetInformationByName(context.Background(), name)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetBookByGenero(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	genero := json_map["genero"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	result, err := service.GetInstanceBook().GetInformationByGenero(context.Background(), genero)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetBookByAutor(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	autor := json_map["autor"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	result, err := service.GetInstanceBook().GetInformationByAutor(context.Background(), autor)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetBookByPriority(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	priority := json_map["priority"].(string)

	result, err := service.GetInstanceBook().GetInformationByPriority(context.Background(), priority)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
