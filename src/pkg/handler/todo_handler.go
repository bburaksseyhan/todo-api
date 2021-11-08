package handler

import (
	"net/http"

	"github.com/bburaksseyhan/todo-api/src/cmd/utils"
	"github.com/bburaksseyhan/todo-api/src/pkg/model"
	database "github.com/bburaksseyhan/todo-api/src/pkg/repository/db"
	log "github.com/labstack/gommon/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoHandler interface {
	CreateTodo(*gin.Context)
	GetTodoById(*gin.Context)

	HealthCheck(*gin.Context)
}

type todoHandler struct {
	repo database.TodoRepository
}

func NewTodoHandler(repo *database.TodoRepository) TodoHandler {
	return &todoHandler{repo: *repo}
}

func (t *todoHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (t *todoHandler) CreateTodo(c *gin.Context) {
	var todo model.Todo

	c.BindJSON(&todo)

	todo.Id = uuid.New().String()
	log.Info("Data is ", todo)
	data, err := t.repo.Save(todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequestError("insert operation failed!", err))
	}

	c.JSON(http.StatusCreated, gin.H{"todo": data})
}

func (t *todoHandler) GetTodoById(c *gin.Context) {

	id := c.Param("id")

	todo, err := t.repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequestError("todo not found", err))
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}
