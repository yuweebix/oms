include .env
ENV_FILE := .env

# СБОРКА
build:
	docker-compose build

# МИГРАЦИЯ
migrate:
	docker-compose --env-file $(ENV_FILE) run --rm app sh -c 'goose -dir db/migrations postgres "$(DATABASE_URL)" up'

# ЗАПУСК
up:
	docker-compose --env-file $(ENV_FILE) up -d
up-app:
	docker-compose --env-file $(ENV_FILE) up -d app
up-db:
	docker-compose --env-file $(ENV_FILE) up -d db

# ОСТАНОВКА
down:
	docker-compose --env-file $(ENV_FILE) down
down-app:
	docker-compose --env-file $(ENV_FILE) down app
down-db:
	docker-compose --env-file $(ENV_FILE) down db

# ПРОЦЕССЫ
cli: # утилита
	docker exec -it $(shell docker-compose ps -q app) sh -c './main'
sql: # sql-клиент
	docker exec -it $(shell docker-compose ps -q db) sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)'
log: # логгер
	docker exec -it $(shell docker-compose ps -q app) tail -f ./log.txt

# SHELL
shell-app:
	docker exec -it $(shell docker-compose ps -q app) sh
shell-db:
	docker exec -it $(shell docker-compose ps -q db) sh
