include .env

# Atlasコマンド
.PHONY: migrate
migrate:
	@echo "Applying schema changes to the local environment..."
	atlas schema apply --env local


# Goコマンド
.PHONY: build run
build:
	@echo "Building Go App..."
	go build -o main main.go
run: 
	@echo "Preparing runing Go App"
	go run main.go


# Dockerコマンド
.PHONY: up down psql
up:
	@echo "Preparing containers..."
	docker compose up -d
down:
	@echo "Closing containers..."
	docker compose down
psql:
	@echo "Preparing Postgres..."
	docker compose exec db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)