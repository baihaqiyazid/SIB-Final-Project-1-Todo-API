package services

import (
	"todo-api/models"
	"todo-api/web"

	"github.com/gin-gonic/gin"
)


type TodoService interface {
	CreateTodo(ctx *gin.Context, request web.TodoPayload) (models.Todo, error)
	GetAll(ctx *gin.Context) []models.Todo
	UpdateTodo(ctx *gin.Context, request web.TodoPayload) (models.Todo, error)
	GetTodoById(ctx *gin.Context, id int) (models.Todo, error)
	DeleteTodo(ctx *gin.Context, id int) error
}