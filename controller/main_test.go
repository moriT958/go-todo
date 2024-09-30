package controller_test

import (
	"go-todo/controller"
	"go-todo/controller/tests"
	"testing"
)

var TC *controller.TodoController

func TestMain(m *testing.M) {
	s := tests.NewServiceMock()
	TC = controller.NewTodoController(s)

	m.Run()
}
