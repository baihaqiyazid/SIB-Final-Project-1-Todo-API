package repository

import (
	"todo-api/helper"
	"todo-api/models"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{
		db: db,
	}
}

// Create implements TodoRepository.
func (repository *TodoRepositoryImpl) Create(todo models.Todo) (models.Todo, error) {
	err := repository.db.Create(&todo).Error
	helper.LogIfError(err)

	return todo, nil
}

// Delete implements TodoRepository.
func (repository *TodoRepositoryImpl) Delete(id int) {
	var todo models.Todo
	repository.db.Where("id", id).Delete(&todo)
}

// GetAll implements TodoRepository.
func (repository *TodoRepositoryImpl) GetAll() []models.Todo {
	var todos []models.Todo

	err := repository.db.Find(&todos).Error
	helper.LogIfError(err)

	return todos
}

// GetTodoById implements TodoRepository.
func (repository *TodoRepositoryImpl) GetTodoById(id int) (models.Todo, error) {
	var todo models.Todo
	err := repository.db.Where("id", id).First(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

// Updatetodo implements TodoRepository.
func (repository *TodoRepositoryImpl) UpdateTodo(todo models.Todo, id int) (models.Todo, error) {
	err := repository.db.Where("id = ?", id).Updates(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

