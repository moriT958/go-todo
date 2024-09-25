package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type PostgresManager struct {
	db *sql.DB
}

func NewPostgresManager() (*PostgresManager, error) {
	loadEnv()
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbConn := fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUser, dbPass, dbName)

	pm := new(PostgresManager)

	db, err := sql.Open("postgres", dbConn)
	pm.db = db

	return pm, err
}

func (pm *PostgresManager) Check() error {
	err := pm.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (pm *PostgresManager) Close() error {
	err := pm.db.Close()
	return err
}

func (pm *PostgresManager) Migrate(query string, args ...any) error {
	if _, err := pm.db.Exec(query, args...); err != nil {
		return err
	}

	log.Println("Successfly migrate table!")
	return nil
}
