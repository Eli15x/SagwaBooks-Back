package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/SagwaBooks-Back/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/sgumirov/go-cards-validation"
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
	userId := json_map["userId"].(string)
	mounth := json_map["mounth"].(string)
	year := json_map["year"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Create Card Error: name not find")
		return
	}

	if numero == "" {
		c.String(400, "Create Card Error: numero not find")
		return
	}

	if mounth == "" {
		c.String(400, "Validate Card Error: data not find")
		return
	}

	if year == "" {
		c.String(400, "Validate Card Error: data not find")
		return
	}

	if userId == "" {
		c.String(400, "Create userId Error: userId not find")
		return
	}

	err = service.GetInstanceCard().CreateNewCard(context.Background(), userId, name, numero, mounth, year)
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
	userId := json_map["userId"].(string)
	mounth := json_map["mounth"].(string)
	year := json_map["year"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Edit Card Error: name not find")
		return
	}

	if numero == "" {
		c.String(400, "Edit Card Error: numero not find")
		return
	}

	if mounth == "" {
		c.String(400, "Validate Card Error: data not find")
		return
	}

	if year == "" {
		c.String(400, "Validate Card Error: data not find")
		return
	}

	if userId == "" {
		c.String(400, "Edit Card Error: userId not find")
		return
	}

	err = service.GetInstanceCard().EditCard(context.Background(), userId, name, numero, mounth, year)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func ValidatedCard(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	numero := json_map["numero"].(string)
	mounth := json_map["mounth"].(string)
	year := json_map["year"].(string)
	userId := json_map["userId"].(string)
	cvv := json_map["cvv"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Edit Card Error: name not find")
		return
	}

	if numero == "" {
		c.String(400, "Validate Card Error: numero not find")
		return
	}

	if mounth == "" {
		c.String(400, "Validate Card Error: data not find")
		return
	}

	if year == "" {
		c.String(400, "Validate Card Error: data not find")
		return
	}

	if userId == "" {
		c.String(400, "Validate Card Error: userId not find")
		return
	}

	if cvv == "" {
		c.String(400, "Validate Card Error: cvv not find")
		return
	}

	// Initialize a new card:
	card := &cards.Card{Number: numero, Cvv: cvv, Month: mounth, Year: year}

	// Validate the card's number (without capturing)
	err = card.Validate() // will return an error due to not allowing test cards

	if err != nil {
		c.String(400, "Validate Card Error: validate card error")
	}
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
