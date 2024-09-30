package service

import "go-todo/models"

type ServiceInterface interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
	ReadTodos(page int) ([]models.Todo, error)
	ReadTodoByID(id int) (models.Todo, error)
	CompleteTodo(id int) (models.Todo, error)
	DeleteTodo(id int) (models.Todo, error)
}
