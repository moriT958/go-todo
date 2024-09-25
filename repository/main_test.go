package repository_test

import (
	"go-todo/repository"
	"os"
	"testing"
)

var TestPM *repository.PostgresManager

func TestMain(m *testing.M) {
	// Get connection to DB
	TestPM, err := repository.NewPostgresManager()
	if err != nil {
		os.Exit(1)
	}
	defer TestPM.Close()

	err = setup(TestPM)
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	err = tearDown(TestPM)
	if err != nil {
		os.Exit(1)
	}
}

func setup(tpm *repository.PostgresManager) error {

	// DB connection check
	err := tpm.Check()
	if err != nil {
		return err
	}

	// migrate table
	if err = tpm.Migrate(`
		CREATE TABLE IF NOT EXISTS testDB (
			id SERIAL PRIMARY KEY,
			task VARCHAR(20) UNIQUE NOT NULL,
			done BOOLEAN NOT NULL DEFAULT false,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			completed_at TIMESTAMP DEFAULT NULL
		);
	`); err != nil {
		return err
	}

	return nil
}

func tearDown(tpm *repository.PostgresManager) error {
	if err := tpm.Migrate(`
		DROP TABLE testDB;
	`); err != nil {
		return err
	}
	return nil
}
