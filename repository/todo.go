package repository

import (
	"database/sql"
	"errors"
	"go-todo/models"
)

const todoNumPerPage = 10

// DBから得たデータを構造体に変換して返す。
func CreateTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	// (注)returningはrowを返す
	const query = `INSERT INTO todos (task) VALUES ($1) RETURNING id, done, created_at;`

	var newTodo models.Todo
	newTodo.Task = todo.Task

	row := db.QueryRow(query, newTodo.Task)
	if err := row.Err(); err != nil {
		return models.Todo{}, err
	}
	if err := row.Scan(&newTodo.TodoID, &newTodo.Done, &newTodo.CreatedAt); err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

func ReadTodos(db *sql.DB, page int) ([]models.Todo, error) {
	const query = `SELECT * FROM todos LIMIT $1 OFFSET $2;`
	if page <= 0 {
		err := errors.New("ページは1以上からです")
		return []models.Todo{}, err
	}

	var createdTime sql.NullTime

	rows, err := db.Query(query, todoNumPerPage, (page-1)*todoNumPerPage)
	if err != nil {
		return nil, err // ※スライスのゼロ値はnil
	}

	todoArray := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.TodoID, &todo.Task, &todo.Done, &createdTime)

		if createdTime.Valid {
			todo.CreatedAt = createdTime.Time
		}

		todoArray = append(todoArray, todo)
	}

	return todoArray, nil
}

func ReadTodoByID(db *sql.DB, id int) (models.Todo, error) {
	const query = `SELECT * FROM todos WHERE id = $1;`

	var gotTodo models.Todo
	var createdTime sql.NullTime // postgres側の値がNULLの可能性があるため

	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return models.Todo{}, err
	}

	if err := row.Scan(&gotTodo.TodoID, &gotTodo.Task, &gotTodo.Done, &createdTime); err != nil {
		return models.Todo{}, err
	}

	// sql側でNULLじゃなかったらgoの構造体に代入
	if createdTime.Valid {
		gotTodo.CreatedAt = createdTime.Time
	}

	return gotTodo, nil
}

func CompleteTodo(db *sql.DB, id int) (models.Todo, error) {
	const query = `UPDATE todos SET done = true WHERE id = $1 RETURNING id;`

	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return models.Todo{}, err
	}

	var completedTodo models.Todo
	completedTodo.TodoID = id
	if err := row.Scan(&completedTodo.Task); err != nil {
		return models.Todo{}, err
	}

	return completedTodo, nil
}

// func DeleteTodo(db *sql.DB, id int) error {

// }
