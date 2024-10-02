package controller

import (
	"encoding/json"
	"go-todo/controller/service"
	"go-todo/models"
	"go-todo/schemas"
	"log"
	"net/http"
	"strconv"
)

// DI (controller -> seviceへの依存。serviceはインターフェースとして与える。)
type TodoController struct {
	service service.ServiceInterface
}

func NewTodoController(s service.ServiceInterface) *TodoController {
	return &TodoController{service: s}
}

func (c *TodoController) HelloHandler(w http.ResponseWriter, _ *http.Request) {
	res := map[string]string{
		"Hello,": "World!",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

func (c *TodoController) PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var reqTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&reqTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("decode error at PostTodoHandler")
		return
	}

	createdTodo, err := c.service.CreateTodo(reqTodo)
	if err != nil {
		log.Println("create fail error at PostTodoHandler")
		return
	}

	// レスポンススキーマに変換
	var res schemas.PostTodoResponse
	res.TodoID, res.Task = createdTodo.TodoID, createdTodo.Task

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func (c *TodoController) GetTodoList(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()

	page := 1

	// pageパラメータが存在し、かつ値が有効である場合
	if pageStr, ok := queryMap["page"]; len(pageStr) > 0 && ok {
		if pageInt, err := strconv.Atoi(pageStr[0]); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("query parse error at GetTodoList")
			return
		} else {
			page = pageInt
		}
	}

	todoList, err := c.service.ReadTodos(page)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	// レスポンススキーマに変換
	var res schemas.GetTodoListResponse
	for _, todo := range todoList {
		res.Todos = append(res.Todos, schemas.Todo{
			TodoID:    todo.TodoID,
			Task:      todo.Task,
			Done:      todo.Done,
			CreatedAt: todo.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

// TODO: 残りのハンドラ実装する
