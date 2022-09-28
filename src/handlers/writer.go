package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Eli15x/SagwaBooks-Back/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateWriter(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	birthDate := json_map["birthDate"].(string)
	email := json_map["email"].(string)
	name := json_map["name"].(string)
	password := json_map["password"].(string)
	city := json_map["city"].(string)
	rg := json_map["rg"].(string)
	cpf := json_map["cpf"].(string)
	telefone := json_map["telefone"].(string)
	image := json_map["image"].(string)

	if birthDate == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: birthDate not find")
		return
	}

	if email == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: email not find")
		return
	}

	if name == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: name not find")
		return
	}

	if password == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: password not find")
		return
	}

	if city == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: city not find")
		return
	}

	if rg == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: rg not find")
		return
	}

	if cpf == "" {
		c.String(http.StatusBadRequest, "Create cpf Error: cpf not find")
		return
	}

	if telefone == "" {
		c.String(http.StatusBadRequest, "Create telefone Error: telefone not find")
		return
	}
	if image == "" {
		c.String(http.StatusBadRequest, "Create image Error: image not find")
		return
	}

	_, err = service.GetInstanceWriter().CreateNewWriter(context.Background(), email, password, birthDate, name, city, rg, cpf, telefone, image)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func EditWriter(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	writerId := json_map["writerId"].(string)
	birthDate := json_map["birthDate"].(string)
	email := json_map["email"].(string)
	name := json_map["name"].(string)
	password := json_map["password"].(string)
	city := json_map["city"].(string)
	rg := json_map["rg"].(string)
	cpf := json_map["cpf"].(string)
	telefone := json_map["telefone"].(string)
	image := json_map["image"].(string)

	if writerId == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: writerId not find")
		return
	}
	if birthDate == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: birthDate not find")
		return
	}

	if email == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: email not find")
		return
	}

	if name == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: name not find")
		return
	}

	if password == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: password not find")
		return
	}

	if city == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: city not find")
		return
	}

	if rg == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: rg not find")
		return
	}

	if cpf == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: cpf not find")
		return
	}

	if telefone == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: telefone not find")
		return
	}
	if image == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: image not find")
		return
	}

	err = service.GetInstanceWriter().EditWriter(context.Background(), writerId, email, password, birthDate, name, city, rg, cpf, telefone, image)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteWriter(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	WriterId := json_map["writerId"].(string)

	if WriterId == "" {
		c.String(http.StatusBadRequest, "Delete Writer Error: userId not find")
		return
	}

	err = service.GetInstanceWriter().DeleteWriter(context.Background(), WriterId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetInformationWriter(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	WriterId := json_map["writerId"].(string)

	if WriterId == "" {
		c.String(http.StatusBadRequest, "Delete Writer Error: userId not find")
		return
	}

	result, err := service.GetInstanceWriter().GetInformationWriter(context.Background(), WriterId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetInformationWriters(c *gin.Context) {

	result, err := service.GetInstanceWriter().GetInformationWriters(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
