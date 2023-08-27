include .env

.SILENT:
.DEFAULT_GOAL = run

CMD_UP = docker-compose up --remove-orphans
CMD_DOWN = docker-compose down --remove-orphans

MIGRATION_DIR = ./scripts/migrations/

POSTGRES_URL = postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(LOCAL_DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	$(CMD_UP)

rebuild: build
	$(CMD_UP) --build

up-postgres:
	$(CMD_UP) postgres

stop:
	$(CMD_DOWN)

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) 'user_segmentation'

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(POSTGRES_URL) up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database $(POSTGRES_URL) down

lint:
	golangci-lint run

swag:
	swag init -g ./internal/app/app.go

clean:
	rm -rf ./.bin

.PHONY: build run rebuild up-postgres stop migrate-create migrate-down migrate-up clean