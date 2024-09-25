package main

import (
	"encoding/json"
	"fmt"
	"go-todo/repository"
	"io"
	"log"
	"net/http"
	"os"

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

	// Get connection to DB
	pm, err := repository.NewPostgresManager()
	if err != nil {
		log.Fatal(err)
	}
	defer pm.Close()

	// DB connection check
	err = pm.Check()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("DB connection is alive")
	}

	// migrate table
	f, err := os.Open("./migrations/1_create_todo_table.sql")
	if err != nil {
		fmt.Println("Error opening migration file:", err)
		return
	}
	defer f.Close()
	b, _ := io.ReadAll(f)
	if err = pm.Migrate(string(b)); err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
