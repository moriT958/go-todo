package fixture

import (
	"database/sql"
)

const dbConn = "postgres://postgres:postgres@127.0.0.1:5432/mydb?sslmode=disable"

type Fixture struct {
	DB *sql.DB
	Tx *sql.Tx
}

func NewFixture() (*Fixture, error) {
	// テスト用DBとの接続
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	// テストトランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return &Fixture{
		db,
		tx,
	}, nil
}

func (f *Fixture) Setup() error {
	// テストテーブルの作成
	if _, err := f.Tx.Exec(`
		DROP TABLE IF EXISTS todos;
		CREATE TABLE "todos" (
			"id" bigserial NOT NULL,
			"task" character varying NOT NULL,
			"done" boolean NOT NULL DEFAULT false,
			"created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY ("id"),
			CONSTRAINT "todos_task_key" UNIQUE ("task")
		);
	`); err != nil {
		return err
	}

	// テストデータの挿入
	for _, d := range TodoTestData {
		if _, err := f.Tx.Exec(`INSERT INTO todos (task) VALUES ($1);`, d.Task); err != nil {
			return err
		}
	}

	return nil
}

func (f *Fixture) Teardown() error {
	// テストテーブルのクリーンアップ
	if _, err := f.Tx.Exec("DROP TABLE IF EXISTS todos;"); err != nil {
		return err
	}

	// ロールバック
	if err := f.Tx.Rollback(); err != nil {
		return err
	}

	// テストDBとの接続を閉じる
	if err := f.DB.Close(); err != nil {
		return err
	}

	return nil
}
