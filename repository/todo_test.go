package repository_test

import (
	"go-todo/models"
	"go-todo/repository"
	"go-todo/repository/testdata"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	testTodo := models.Todo{
		Task: "test1",
	}

	const expectedTodoID = 4

	newTodo, err := repository.CreateTodo(testDB, testTodo)
	if err != nil {
		t.Fatal(err)
	}

	if newTodo.TodoID != expectedTodoID {
		t.Errorf("Want Todo Task %d, but Got %d\n", expectedTodoID, newTodo.TodoID)
	}

	t.Cleanup(func() {
		const query = "DELETE FROM todos WHERE task = $1;"
		testDB.Exec(query, testTodo.Task)
	})
}

func TestReadTodos(t *testing.T) {
	expectedNum := len(testdata.TodoTestData)

	got, err := repository.ReadTodos(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("Want %d, but Got %d\n", expectedNum, num)
	}
}

func TestReadTodoByID(t *testing.T) {
	// テーブルドリブンテスト
	tests := []struct {
		testTitle string
		expected  models.Todo
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.TodoTestData[0],
		},
		{
			testTitle: "subtest2",
			expected:  testdata.TodoTestData[1],
		},
		{
			testTitle: "subtest3",
			expected:  testdata.TodoTestData[2],
		},
	}

	for i, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repository.ReadTodoByID(testDB, i+1)
			if err != nil {
				t.Fatal(err)
			}

			if got.Task != test.expected.Task {
				t.Errorf("Task: Want %s, but Got %s\n", test.expected.Task, got.Task)
			}
		})
	}
}
