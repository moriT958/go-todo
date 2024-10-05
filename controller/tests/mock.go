package tests

import (
	"errors"
	"go-todo/models"
)

type ServicMock struct{}

func NewServiceMock() *ServicMock {
	return &ServicMock{}
}

func (sm ServicMock) CreateTodo(todo models.Todo) (models.Todo, error) {
	if todo.Task == "" {
		return models.Todo{}, errors.New("validation error(task is empty)")
	}

	return TodoTestData[0], nil
}

func (sm ServicMock) ReadTodos(page int) ([]models.Todo, error) {
	return TodoTestData, nil
}

func (sm ServicMock) ReadTodoByID(id int) (models.Todo, error) {
	if !(1 <= id && id <= 3) {
		return models.Todo{}, errors.New("at Service ReadTodoByID, no data found")
	}

	return TodoTestData[id-1], nil
}

func (sm ServicMock) CompleteTodo(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (sm ServicMock) DeleteTodo(id int) (models.Todo, error) {
	return models.Todo{}, nil
}
