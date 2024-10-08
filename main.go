package main

import (
	"database/sql"
	"go-todo/controller"
	"go-todo/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	loadEnv()
	dsn := os.Getenv("DATABASE_URL")

	// Get connection to DB
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Connect Database!")
	}
	defer db.Close()

	// DB connection check
	if err := db.Ping(); err != nil {
		log.Println(err)
	} else {
		log.Println("Your Database is Alive!")
	}

	s := service.NewService(db)
	tc := controller.NewTodoController(s)

	// TODO: ルータのまとめるリファクタ
	r := mux.NewRouter()
	r.HandleFunc("/api", tc.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", tc.PostTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/todo", tc.GetTodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id:[0-9+]}", tc.GetTodoByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id:[0-9+]}", tc.CompleteTodoHandler).Methods(http.MethodPatch)
	r.HandleFunc("/api/todo/{id:[0-9+]}", tc.DeleteTodoHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))

	// TODO: バリデーションミドルウェア実装
	// TODO: エラーまとめるリファクタ
}

func loadEnv() error {
	err := godotenv.Load(".env")
	return err
}
