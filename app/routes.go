package app

import (
	"todo-api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func Route(todoController controller.TodoController){
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	
	router.POST("/api/v1/todos", todoController.CreateTodo)
	router.GET("/api/v1/todos", todoController.GetAllTodo)
	router.PUT("/api/v1/todos/:id", todoController.UpdateTodo)
	router.DELETE("/api/v1/todos/:id", todoController.DeleteTodo)
	router.GET("/api/v1/todos/:id", todoController.GetTodoById)

	router.Run(":4001")
}