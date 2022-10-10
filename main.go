package main

import (
	//"time"
	//"context"

	"context"
	"fmt"
	"time"

	"github.com/Eli15x/SagwaBooks-Back/src/client"
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
	if err := client.GetInstance().Initialize(ctx); err != nil {
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
	router.POST("/user/edit", handlers.EditUser)
	router.POST("/user/delete", handlers.DeleteUser)
	router.POST("/writer/create", handlers.CreateWriter)
	router.POST("/writer/get", handlers.GetInformationWriter)
	router.GET("/writer/getAll", handlers.GetInformationWriters)
	router.POST("/writer/edit", handlers.EditWriter)
	router.POST("/writer/delete", handlers.DeleteWriter)
	router.POST("/getInformation", handlers.GetInformationByUserId)
	router.POST("/book/create", handlers.CreateBook)
	router.POST("/book/edit", handlers.EditBook)
	router.POST("/book/delete", handlers.DeleteBook)
	router.POST("/book/name", handlers.GetBookByName)
	router.POST("/book/autor", handlers.GetBookByAutor)
	router.POST("/book/genero", handlers.GetBookByGenero)
	router.POST("/book/priority", handlers.GetBookByPriority)
	router.POST("/card/create", handlers.CreateCard)
	router.POST("/card/edit", handlers.EditCard)
	router.POST("/card/delete", handlers.DeleteCard)
	router.POST("/card/user", handlers.GetCardsByUserId)
	router.POST("/card/validate", handlers.ValidatedCard)

	router.Run(":1323")
}
