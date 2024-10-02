package schemas

import "time"

type Todo struct {
	TodoID    int       `json:"todo_id"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type PostTodoResponse struct {
	TodoID int    `json:"todo_id"`
	Task   string `json:"task"`
}

type GetTodoListResponse struct {
	Todos []Todo `json:"data"`
}

// TODO: そのほかのスキーマも設定する。
