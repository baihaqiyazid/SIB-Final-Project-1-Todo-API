package services

import (
	"strconv"
	"time"
	"todo-api/helper"
	"todo-api/models"
	"todo-api/repository"
	"todo-api/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoServiceImpl struct {
	DB             *gorm.DB
	TodoRepository repository.TodoRepository
}

func NewTodoService(db *gorm.DB, todoRepository repository.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{
		DB:             db,
		TodoRepository: todoRepository,
	}
}

func (service *TodoServiceImpl) CreateTodo(ctx *gin.Context, request web.TodoPayload) (models.Todo, error) {
	var todo models.Todo

	todo.Text = request.Text
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	data, err := service.TodoRepository.Create(todo)
	helper.LogIfError(err)

	return data, nil
}

func (service *TodoServiceImpl) GetAll(ctx *gin.Context) []models.Todo {
	data := service.TodoRepository.GetAll()

	return data
}

func (service *TodoServiceImpl) UpdateTodo(ctx *gin.Context, request web.TodoPayload) (models.Todo, error) {
	var todo models.Todo

	id := ctx.Param("id")
	todoId, _ := strconv.Atoi(id)

	todo.Text = request.Text
	todo.UpdatedAt = time.Now()

	_, err := service.TodoRepository.UpdateTodo(todo, todoId)
	if err != nil {
		return models.Todo{}, err
	}

	data, err := service.TodoRepository.GetTodoById(todoId)
	if err != nil {
		return models.Todo{}, err
	}

	return data, nil
}

func (service *TodoServiceImpl) GetTodoById(ctx *gin.Context, id int) (models.Todo, error) {
	data, err := service.TodoRepository.GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}

	return data, nil
}

func (service *TodoServiceImpl) DeleteTodo(ctx *gin.Context, id int) error {

	service.TodoRepository.Delete(id)

	return nil
}
