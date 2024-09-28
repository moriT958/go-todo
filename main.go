package main

import (
	"database/sql"
	"encoding/json"
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

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func loadEnv() error {
	err := godotenv.Load(".env")
	return err
}
