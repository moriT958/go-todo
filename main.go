package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-todo/migrations"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	res := map[string]string{
		"Hello,": "World!",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

func main() {
	loadEnv()
	dbUsr := os.Getenv("POSTGRES_USER")
	dbPwd := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbConn := fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUsr, dbPwd, dbName)

	// Get connection to DB
	db, err := sql.Open("postgres", dbConn)
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

	// DB Migration
	if err := migrations.Migrate(db, "migrations/0_create_todo_table.sql"); err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func loadEnv() error {
	err := godotenv.Load(".env")
	return err
}
