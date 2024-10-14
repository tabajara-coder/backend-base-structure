ifneq (,$(wildcard ./.env))
    include .env
    export
endif


ifeq ($(OS),Windows_NT)
	MAIN_PATH = tmp\bin\main.exe
else
	MAIN_PATH = tmp/bin/main
endif

DB_DBSTRING := $(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

init: start-compose db-up server

start-docker:
	@echo "Checking if Docker is running..."
	@if ! docker info > /dev/null 2>&1; then \
		echo "Docker is not running. Please start Docker manually."; \
		exit 1; \
	else \
		echo "Docker is running."; \
	fi

start-compose: start-docker
	@echo "Starting Docker Compose..."
	@docker-compose up -d

server:
	@go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build --tags dev -o ${MAIN_PATH} ./cmd/" \
		--build.bin "${MAIN_PATH}" \
		--build.delay "100" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true \
		--screen.clear_on_rebuild true \
		--log.main_only true

build:
	@go build -o bin/app_prod cmd/app/main.go
	@echo "compiled you application with all its assets to a single binary => bin/app_prod"

db-seed:
	@go run cmd/scripts/seed/main.go

db-down:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_DBSTRING) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) down

db-up:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_DBSTRING) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) up

migration:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_DBSTRING) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) create $(filter-out $@,$(MAKECMDGOALS)) sql
