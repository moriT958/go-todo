package repository_test

import (
	"go-todo/models"
	"go-todo/repository"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	testTodo := models.Todo{
		Task: "test todo1",
	}

	const expectedID = 1
	const expectedTask = "test todo1"

	var newTodo models.Todo
	newTodo, err := repository.CreateTodo(*TestPM, testTodo)
	if err != nil {
		t.Error(err)
	}
	if newTodo.TodoID != expectedID {
		t.Errorf("new todo id is expected %d but got %d\n", expectedID, newTodo.TodoID)
	}
	if newTodo.Task != expectedTask {
		t.Errorf("new todo task is expected %s but got %s\n", expectedTask, newTodo.Task)
	}

	t.Cleanup(func() {
		const query = `
			DELETE FROM testDB
			where task = $1;
		`
		TestPM.Migrate(query, testTodo.Task)
	})
}
