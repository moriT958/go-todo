package repository

import (
	"database/sql"
	"errors"
	"go-todo/models"
)

const todoNumPerPage = 10

func CreateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	// (注)returningはrowを返す
	const query = `INSERT INTO todos (task) VALUES ($1) RETURNING id;`

	var newTodo models.Todo
	newTodo.Task = todo.Task

	row := db.QueryRow(query, newTodo.Task)
	if err := row.Err(); err != nil {
		return models.Todo{}, err
	}
	if err := row.Scan(&newTodo.TodoID); err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil // idとtaskのみ返す
}

func ReadTodos(db *sql.DB, page int) ([]models.Todo, error) {
	const query = `SELECT * FROM todos LIMIT $1 OFFSET $2;`
	if page <= 0 {
		err := errors.New("ページは1以上からです")
		return []models.Todo{}, err
	}

	rows, err := db.Query(query, todoNumPerPage, (page-1)*todoNumPerPage)
	if err != nil {
		return nil, err // ※スライスのゼロ値はnil
	}

	todoArray := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.TodoID, &todo.Task, &todo.Done, &todo.CreatedAt, &todo.CompletedAt)

		todoArray = append(todoArray, todo)
	}

	return todoArray, nil
}

func ReadTodoByID(db *sql.DB, id int) (models.Todo, error) {
	const query = `SELECT * FROM todos WHERE id = $1;`

	var gotTodo models.Todo
	var createdTime, completedTime sql.NullTime // postgres側で設定されたNULL値がGoで扱えない可能性があるため

	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return models.Todo{}, err
	}

	if err := row.Scan(&gotTodo.TodoID, &gotTodo.Task, &gotTodo.Done, &createdTime, &completedTime); err != nil {
		return models.Todo{}, err
	}

	// sql側で設定された時間がNULLじゃなかったらgoの構造体に代入
	if createdTime.Valid {
		gotTodo.CreatedAt = createdTime.Time
	}
	if completedTime.Valid {
		gotTodo.CompletedAt = completedTime.Time
	}

	return gotTodo, nil
}

// func CompleteTodo(db *sql.DB, id int) (models.Todo, error) {
// 	const query = `UPDATE `

// 	tx, err := db.Begin()
// 	if err != nil {
// 		return models.Todo{}, err
// 	}

// }
