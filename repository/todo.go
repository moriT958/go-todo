package repository

import "go-todo/models"

func CreateTodo(pm PostgresManager, todo models.Todo) (models.Todo, error) {

	const query = `
		INSERT INTO todos (task)
		VALUES ($1);
	`
	var newTodo models.Todo
	newTodo.Task = todo.Task

	row := pm.db.QueryRow(query, newTodo.Task)
	if err := row.Scan(&newTodo.TodoID); err != nil {
		return models.Todo{}, err
	}
	return newTodo, nil
}
