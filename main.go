package main

import (
	"todo-api/app"
	"todo-api/controller"
	"todo-api/repository"
	_ "todo-api/docs"
	"todo-api/services"
)

// @title 	Todos Service API
// @version	1.0
// @description A Todo service API in Go using Gin framework

// @host 	localhost:4001
// @BasePath /api/v1
func main()  {
	app.StartDB()

	db := app.GetDB()

	todoRepository := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(db, todoRepository)
	todoController := controller.NewTodoController(todoService)

	app.Route(todoController)
}