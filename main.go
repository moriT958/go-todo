package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-todo/models"
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

	loadEnv() // Load env-vars from .env file
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbConn := fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPass, dbName)

	// Get connection to DB
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Println("fail to get db connection...")
		return
	}
	defer db.Close()

	// DB connection check
	err = db.Ping()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("DB connection is alive")
	}

	err = models.InitDB(db)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("fail to load envfile...")
	}
}
