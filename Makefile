include .env

.SILENT:
.PHONY:  up-postgres stop migrate-create migrate-down migrate-up

CMD_UP = docker-compose up --remove-orphans
CMD_DOWN = docker-compose down

MIGRATION_DIR = ./scripts/migrations/

POSTGRES_URL = postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(LOCAL_DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

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