package tests

import "go-todo/models"

type ServicMock struct{}

func NewServiceMock() *ServicMock {
	return &ServicMock{}
}

func (sm ServicMock) CreateTodo(todo models.Todo) (models.Todo, error) {
	return TodoTestData[0], nil
}

func (sm ServicMock) ReadTodos(page int) ([]models.Todo, error) {
	return TodoTestData, nil
}

func (sm ServicMock) ReadTodoByID(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (sm ServicMock) CompleteTodo(id int) (models.Todo, error) {
	return models.Todo{}, nil
}

func (sm ServicMock) DeleteTodo(id int) (models.Todo, error) {
	return models.Todo{}, nil
}
