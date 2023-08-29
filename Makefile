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

test:
	go test -coverprofile=cover.out -v ./...
	make --silent test-coverage

test-coverage:
	go tool cover -func=cover.out | grep "total"

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) 'user_segmentation'

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(POSTGRES_URL) up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database $(POSTGRES_URL) down

lint:
	golangci-lint run

swag:
	 swag init -g internal/app/app.go -o ./docs/swagger/

clean:
	rm -rf ./.bin cover.out

.PHONY: build run rebuild up-postgres stop test test-coverage html-coverage migrate-create migrate-down migrate-up lint swag clean