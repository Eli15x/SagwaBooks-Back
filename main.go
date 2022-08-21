package main

import (
	//"time"
	//"context"

	"context"
	"fmt"
	"time"

	storage "github.com/Eli15x/SagwaBooks-Back/src/client"
	"github.com/Eli15x/SagwaBooks-Back/src/handlers"
	"github.com/bugsnag/bugsnag-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//bugsnag configure
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "3ecac0ed23b7b1f4b863073135c602b8",
		ReleaseStage: "production",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/org/myapp"},
		// more configuration options
	})

	bugsnag.Notify(fmt.Errorf("Test error"))

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	//Connection to Mongo
	if err := storage.GetInstance().Initialize(ctx); err != nil {
		bugsnag.Notify(fmt.Errorf("[MONGO DB - MovieWorkNow] Could not resolve Data access layer. Error:"))
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/login", handlers.ValidateUser)
	router.POST("/cadastro", handlers.CreateUser)
	router.POST("/getInformation", handlers.GetInformationByUserId)
	router.POST("/book/create", handlers.CreateBook)
	router.POST("/book/edit", handlers.EditBook)
	router.POST("/book/delete", handlers.DeleteBook)
	router.POST("/book/name", handlers.GetBookByName)
	router.POST("/book/autor", handlers.GetBookByAutor)
	router.POST("/book/genero", handlers.GetBookByGenero)

	router.Run(":1323")
}