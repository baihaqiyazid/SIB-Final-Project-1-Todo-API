package controller

import "github.com/gin-gonic/gin"

type TodoController interface{
	CreateTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	GetAllTodo(ctx *gin.Context)
	GetTodoById(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}
