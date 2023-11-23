package repository

import "todo-api/models"

type TodoRepository interface{
	Create(todo models.Todo) (models.Todo, error)
	GetAll() []models.Todo
	UpdateTodo(todo models.Todo, id int) (models.Todo, error)
	GetTodoById(id int) (models.Todo, error)
	Delete(id int)
}