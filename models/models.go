package models

import (
	"time"
)

type Todo struct {
	TodoID    int
	Task      string
	Done      bool
	CreatedAt time.Time
}

// TODO: ORM使ってデータベーススキーマ定義する。
