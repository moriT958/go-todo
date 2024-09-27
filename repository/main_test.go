package repository_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var testDB *sql.DB

// テスト用DBの環境変数
var (
	dbUsr  = "postgres"
	dbPwd  = "postgres"
	dbName = "mydb"
	dbConn = fmt.Sprintf("postgres://%s:%s@127.0.0.1:5432/%s?sslmode=disable", dbUsr, dbPwd, dbName)
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

func connectDB() error {
	var err error
	testDB, err = sql.Open("postgres", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {
	f, err := os.Open("./testdata/setup.sql")
	if err != nil {
		return err
	}
	defer f.Close()

	query := make([]byte, 1024)
	count, err := f.Read(query)
	if err != nil {
		return err
	}

	_, err = testDB.Exec(string(query[:count]))
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	f, err := os.Open("./testdata/cleanup.sql")
	if err != nil {
		return err
	}
	defer f.Close()

	query := make([]byte, 1024)
	count, err := f.Read(query)
	if err != nil {
		return err
	}

	_, err = testDB.Exec(string(query[:count]))
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		fmt.Println(err, "at connectDB")
		return err
	}

	// DBのクリーンアップ
	if err := cleanupDB(); err != nil {
		fmt.Println(err, "at cleanupDB")
		return err
	}

	// 初期データの挿入
	if err := setupTestData(); err != nil {
		fmt.Println(err, "at setupTestData")
		return err
	}

	return nil
}

func teardown() {
	// DBのクリーンアップ
	cleanupDB()
	testDB.Close()
}
