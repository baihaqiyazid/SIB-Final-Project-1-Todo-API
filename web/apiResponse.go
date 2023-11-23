package web

import (
	"time"
	_ "todo-api/docs"
)

type Todos struct {
	Id        int       `json:"id" example:"1"`
	Text      string    `json:"text" example:"hellooo, this is my First todo yaww!!"`
	CreatedAt time.Time `json:"created_at" example:"2023-11-13T11:51:04.319747+07:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-11-13T11:51:04.319747+07:00"`
}

// ========================================================
type InternalServerErrorResponse struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"internal server error"`
	Data   string `json:"data" example:"Data not found"`
}

type NotFoundResponse struct {
	Code   int    `json:"code" example:"404"`
	Status string `json:"status" example:"not found"`
	Data   string `json:"data" example:"Data not found"`
}

// ========================================================
type CreateAndUpdateSuccessResponse struct {
	Code   int    `json:"code" example:"200"`
	Status string `json:"status" example:"success"`
	Data   Todos  `json:"data"`
}

type GetAllDataSuccessResponse struct {
	Code   int     `json:"code" example:"200"`
	Status string  `json:"status" example:"success"`
	Data   []Todos `json:"data"`
}

type DeleteSuccessResponse struct {
	Code   int    `json:"code" example:"200"`
	Status string `json:"status" example:"success"`
	Data   string `json:"data" example:"null"`
}
