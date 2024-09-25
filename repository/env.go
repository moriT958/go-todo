package repository

import (
	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load(".env")
	return err
}
