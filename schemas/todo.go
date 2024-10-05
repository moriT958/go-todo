package schemas

import "go-todo/models"

type PostTodoRequest struct {
	Task string `json:"task"`
}

type PostTodoResponse struct {
	TodoID int    `json:"todo_id"`
	Task   string `json:"task"`
}

type GetTodoListResponse struct {
	Todos []models.Todo `json:"data"`
}

type GetTodoByIDResponse struct {
	Todo models.Todo `json:"data"`
}

type CompleteTodoResponse struct {
	TodoID int    `json:"todo_id"`
	Task   string `json:"task"`
}

type DeleteTodoResponse struct {
	TodoID int    `json:"todo_id"`
	Task   string `json:"task"`
}
