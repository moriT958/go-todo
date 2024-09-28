include .env

# Atlasコマンド
.PHONY: apply inspect migrate
apply:
	@echo "Applying schema changes to the local environment..."
	atlas schema apply \
  	--url $(DATABASE_URL) \
  	--to "file://database/schema.sql" \
  	--dev-url "docker://postgres/15/dev?search_path=public"
inspect:
	@echo "Inspecting schema..."
	atlas schema inspect \
	--url $(DATABASE_URL) \
	--format '{{ sql . }}' \
	> database/schema.sql
migrate:
	@echo "Migrating schema..."
	atlas migrate diff diffideal \
	--dir "file://database/migrations" \
	--to "file://database/schema.sql" \
	--dev-url "docker://postgres/15/dev?search_path=public"