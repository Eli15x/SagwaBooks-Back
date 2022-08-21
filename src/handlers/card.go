package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/SagwaBooks-Back/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateCard(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	numero := json_map["numero"].(string)
	data := json_map["data"].(string)
	userId := json_map["userId"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Create Card Error: name not find")
		return
	}

	if numero == "" {
		c.String(400, "Create Card Error: numero not find")
		return
	}

	if data == "" {
		c.String(400, "Create Data Error: data not find")
		return
	}

	if userId == "" {
		c.String(400, "Create userId Error: userId not find")
		return
	}

	err = service.GetInstanceCard().CreateNewCard(context.Background(), name, numero, data, userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func EditCard(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	numero := json_map["numero"].(string)
	data := json_map["data"].(string)
	userId := json_map["userId"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Edit Card Error: name not find")
		return
	}

	if numero == "" {
		c.String(400, "Edit Card Error: numero not find")
		return
	}

	if data == "" {
		c.String(400, "Edit Card Error: data not find")
		return
	}

	if userId == "" {
		c.String(400, "Edit Card Error: userId not find")
		return
	}

	err = service.GetInstanceCard().EditCard(context.Background(), name, numero, data, userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteCard(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	numero := json_map["numero"].(string)

	if numero == "" {
		c.String(http.StatusBadRequest, "Delete Card Error: name not find")
		return
	}

	err = service.GetInstanceCard().DeleteCard(context.Background(), numero)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetCardsByUserId(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string)

	result, err := service.GetInstanceCard().GetInformationByUserId(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
