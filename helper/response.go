package helper

import (
	"todo-api/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	})
}

func ResponseBadRequest(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusBadRequest, web.Response{
		Code:   http.StatusBadRequest,
		Status: "bad request",
		Data: data,
	})
}

func ResponseNotFound(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusNotFound, web.Response{
		Code:   http.StatusNotFound,
		Status: "not found",
		Data: data,
	})
}