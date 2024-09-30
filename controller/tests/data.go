package tests

import (
	"go-todo/models"
	"time"
)

var TodoTestData = []models.Todo{
	{
		TodoID:    1,
		Task:      "test-todo1",
		Done:      false,
		CreatedAt: time.Now(),
	},
	{
		TodoID:    2,
		Task:      "test-todo2",
		Done:      true,
		CreatedAt: time.Now(),
	},
	{
		TodoID:    3,
		Task:      "test-todo3",
		Done:      false,
		CreatedAt: time.Now(),
	},
}
