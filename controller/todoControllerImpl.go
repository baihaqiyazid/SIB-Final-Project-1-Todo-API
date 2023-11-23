package controller

import (
	"todo-api/helper"
	"todo-api/services"
	"todo-api/web"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoControllerImpl struct {
	TodoService services.TodoService
}

func NewTodoController(service services.TodoService) *TodoControllerImpl {
	return &TodoControllerImpl{TodoService: service}
}

// CreateTodos		godoc
// @Summary			Create todos
// @Description		Save todos data in Db.
// @Param			todos body web.TodoPayload true "Create todos"
// @Produce			application/json
// @Tags			todos
// @Success			200 {object} web.CreateAndUpdateSuccessResponse{}
// @Failure 		500 {object} web.InternalServerErrorResponse{}
// @Router			/todos [post]
func (controller *TodoControllerImpl) CreateTodo(ctx *gin.Context) {

	var request web.TodoPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.TodoService.CreateTodo(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err)
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

// UpdateTodos		godoc
// @Summary			Update todos
// @Description		Update todos data save in db.
// @Param			todoId path string true "update todos by id"
// @Param			tags body web.TodoPayload true  "Update todos"
// @Tags			todos
// @Produce			application/json
// @Success			200 {object} web.CreateAndUpdateSuccessResponse{}
// @Failure 		500 {object} web.InternalServerErrorResponse{}
// @Failure 		404 {object} web.NotFoundResponse{}
// @Router			/todos/{Id} [put]
func (controller *TodoControllerImpl) UpdateTodo(ctx *gin.Context) {

	var request web.TodoPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	// Check Todo by Id
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	_,err := controller.TodoService.GetTodoById(ctx, id)
	if err != nil{
		fmt.Println(err)
		helper.ResponseNotFound(ctx, "Data Not Found")
		panic(err)
	}

	data, err := controller.TodoService.UpdateTodo(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err)
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

// UpdateTodos		godoc
// @Summary			Get All todos
// @Description		Get All todos data from db.
// @Tags			todos
// @Produce			application/json
// @Success			200 {object} web.GetAllDataSuccessResponse{}
// @Failure 		500 {object} web.InternalServerErrorResponse{}
// @Router			/todos [get]
func (controller *TodoControllerImpl) GetAllTodo(ctx *gin.Context)  {
	data := controller.TodoService.GetAll(ctx)

	helper.ResponseSuccess(ctx, data)

	return
}

// UpdateTodos		godoc
// @Summary			Delete todos
// @Description		Delete todos data in db.
// @Param			todoId path string true "delete todos by id"
// @Tags			todos
// @Produce			application/json
// @Success			200 {object} web.DeleteSuccessResponse{}
// @Failure 		500 {object} web.InternalServerErrorResponse{}
// @Failure 		404 {object} web.NotFoundResponse{}
// @Router			/todos/{Id} [delete]
func (controller *TodoControllerImpl) DeleteTodo(ctx *gin.Context) {
	// Check Todo by Id
	id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := controller.TodoService.GetTodoById(ctx, id)
	if err != nil{
		fmt.Println(err)
		helper.ResponseNotFound(ctx, "Data Not Found")
		panic(err)
	}
	
	controller.TodoService.DeleteTodo(ctx, id)	

	helper.ResponseSuccess(ctx, nil)

	return
}

// UpdateTodos		godoc
// @Summary			Get todos
// @Description		Get single todos data.
// @Param			todoId path string true "Get todos by id"
// @Tags			todos
// @Produce			application/json
// @Success			200 {object} web.CreateAndUpdateSuccessResponse{}
// @Failure 		500 {object} web.InternalServerErrorResponse{}
// @Failure 		404 {object} web.NotFoundResponse{}
// @Router			/todos/{Id} [get]
func (controller *TodoControllerImpl) GetTodoById(ctx *gin.Context) {
	// Check Todo by Id
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := controller.TodoService.GetTodoById(ctx, id)
	if err != nil{
		fmt.Println(err)
		helper.ResponseNotFound(ctx, "Data Not Found")
		panic(err)
	}

	helper.ResponseSuccess(ctx, data)

	return
}