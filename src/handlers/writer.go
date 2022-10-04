package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Eli15x/SagwaBooks-Back/src/models"
	"github.com/Eli15x/SagwaBooks-Back/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateWriter(c *gin.Context) {

	writer := &models.Writer{}
	err := json.NewDecoder(c.Request.Body).Decode(writer)

	if err != nil {
		c.String(400, "%s", err)
		return
	}
	if writer.BirthDate == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: birthDate not find")
		return
	}

	if writer.Email == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: email not find")
		return
	}

	if writer.Name == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: name not find")
		return
	}

	if writer.PassWord == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: password not find")
		return
	}

	if writer.City == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: city not find")
		return
	}

	if writer.Rg == "" {
		c.String(http.StatusBadRequest, "Create Writer Error: rg not find")
		return
	}

	if writer.Cpf == "" {
		c.String(http.StatusBadRequest, "Create cpf Error: cpf not find")
		return
	}

	if writer.Telefone == "" {
		c.String(http.StatusBadRequest, "Create telefone Error: telefone not find")
		return
	}
	if writer.Image == "" {
		c.String(http.StatusBadRequest, "Create image Error: image not find")
		return
	}

	_, err = service.GetInstanceWriter().CreateNewWriter(context.Background(), writer)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func EditWriter(c *gin.Context) {

	writer := &models.Writer{}
	err := json.NewDecoder(c.Request.Body).Decode(writer)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if writer.WriterId == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: writerId not find")
		return
	}
	if writer.BirthDate == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: birthDate not find")
		return
	}

	if writer.Email == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: email not find")
		return
	}

	if writer.Name == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: name not find")
		return
	}

	if writer.PassWord == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: password not find")
		return
	}

	if writer.City == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: city not find")
		return
	}

	if writer.Rg == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: rg not find")
		return
	}

	if writer.Cpf == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: cpf not find")
		return
	}

	if writer.Telefone == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: telefone not find")
		return
	}
	if writer.Image == "" {
		c.String(http.StatusBadRequest, "Edite Writer Error: image not find")
		return
	}

	err = service.GetInstanceWriter().EditWriter(context.Background(), writer)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteWriter(c *gin.Context) {

	writer := &models.Writer{}
	err := json.NewDecoder(c.Request.Body).Decode(writer)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if writer.WriterId == "" {
		c.String(http.StatusBadRequest, "Delete Writer Error: userId not find")
		return
	}

	err = service.GetInstanceWriter().DeleteWriter(context.Background(), writer.WriterId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetInformationWriter(c *gin.Context) {

	writer := &models.Writer{}
	err := json.NewDecoder(c.Request.Body).Decode(writer)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if writer.WriterId == "" {
		c.String(http.StatusBadRequest, "Delete Writer Error: userId not find")
		return
	}

	result, err := service.GetInstanceWriter().GetInformationWriter(context.Background(), writer.WriterId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetInformationWriters(c *gin.Context) {

	fmt.Println("entrou")
	result, err := service.GetInstanceWriter().GetInformationWriters(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}
