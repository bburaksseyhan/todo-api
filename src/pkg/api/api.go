package api

import (
	"fmt"

	"github.com/bburaksseyhan/todo-api/src/cmd/utils"
	"github.com/bburaksseyhan/todo-api/src/pkg/clients/cassandra"
	"github.com/bburaksseyhan/todo-api/src/pkg/handler"
	"github.com/bburaksseyhan/todo-api/src/pkg/repository/db"
	"github.com/gin-gonic/gin"
)

// Initialize GIN Web Framework and prepare routing, Connect to Cassandra Client, Register to Repository and Handler
func Initialize(config utils.Configuration) {

	// Creates a gin router with default middleware:
	router := gin.Default()

	fmt.Printf("%+v\n", config)

	//register database ,repositories and handlers
	session := cassandra.ConnectDatabase(config.Database.Url, config.Database.Keyspace)

	repository := db.NewTodoRepository(session)
	orderHandler := handler.NewTodoHandler(&repository)

	router.GET("/ping", orderHandler.HealthCheck)

	router.POST("api/v1/todo/", orderHandler.CreateTodo)
	router.GET("api/v1/todo/:id", orderHandler.GetTodoById)

	//run the server :8080
	router.Run(":" + config.Host.Port + "")
}
