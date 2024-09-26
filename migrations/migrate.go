package migrations

import (
	"database/sql"
	"log"
	"os"
)

func Migrate(db *sql.DB, filename string) error {

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	query := make([]byte, 1024)
	count, err := f.Read(query)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(query[:count]))
	if err != nil {
		return err
	}

	log.Println("Migration Successed!!!")
	return nil
}
