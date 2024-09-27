include .env

# Atlasコマンド
.PHONY: migrate
migrate:
	@echo "Applying schema changes to the local environment..."
	atlas schema apply \
  	--url "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?search_path=public&sslmode=disable" \
  	--to "file://database/schema.sql" \
  	--dev-url "docker://postgres/15/dev?search_path=public"


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