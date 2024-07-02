include .env
ENV_FILE := .env

# СБОРКА
build:
	docker-compose build
build-app:
	docker-compose build app
build-test:
	docker-compose build app_test

# МИГРАЦИЯ
migrate:
	docker-compose --env-file $(ENV_FILE) run --rm app sh -c 'goose -dir db/migrations postgres "$(DATABASE_URL)" up'
migrate-test:
	docker-compose --env-file $(ENV_FILE) run --rm app_test sh -c 'goose -dir db/migrations postgres "$(DATABASE_TEST_URL)" up'

demigrate:
	docker-compose --env-file $(ENV_FILE) run --rm app sh -c 'goose -dir db/migrations postgres "$(DATABASE_URL)" down'
demigrate-test:
	docker-compose --env-file $(ENV_FILE) run --rm app_test sh -c 'goose -dir db/migrations postgres "$(DATABASE_TEST_URL)" down'

# БД
create-db:
	docker-compose --env-file $(ENV_FILE) run --rm db sh -c 'psql -h db -U $(POSTGRES_USER) -c "CREATE DATABASE $(POSTGRES_DB);"'
create-db-test:
	docker-compose --env-file $(ENV_FILE) run --rm db_test sh -c 'psql -h db_test -U $(POSTGRES_USER) -c "CREATE DATABASE $(POSTGRES_DB)_test;"'

# ЗАПУСК
up:
	docker-compose --env-file $(ENV_FILE) up -d
up-app:
	docker-compose --env-file $(ENV_FILE) up -d app
up-db:
	docker-compose --env-file $(ENV_FILE) up -d db
up-test:
	docker-compose --env-file $(ENV_FILE) up -d app_test

# ОСТАНОВКА
down:
	docker-compose --env-file $(ENV_FILE) down
down-app:
	docker-compose --env-file $(ENV_FILE) down app
down-db:
	docker-compose --env-file $(ENV_FILE) down db
down-test:
	docker-compose --env-file $(ENV_FILE) down app_test db_test

# ПРОЦЕССЫ
cli: # утилита
	docker exec -it $(shell docker-compose ps -q app) sh -c './main'
cli-test: # утилита тестов
	docker exec -it $(shell docker-compose ps -q app_test) sh -c './main'

sql: # sql-клиент
	docker exec -it $(shell docker-compose ps -q db) sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)'
sql-test: # sql-клиент для тестов
	docker exec -it $(shell docker-compose ps -q db_test) sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)_test'

log: # логгер
	docker exec -it $(shell docker-compose ps -q app) tail -f ./log.txt
log-test: # логгер для тестов
	docker exec -it $(shell docker-compose ps -q app_test) tail -f ./log.txt

# SHELL
shell-app:
	docker exec -it $(shell docker-compose ps -q app) sh
shell-db:
	docker exec -it $(shell docker-compose ps -q db) sh
shell-app-test:
	docker exec -it $(shell docker-compose ps -q app_test) sh
shell-db-test:
	docker exec -it $(shell docker-compose ps -q db_test) sh

# МОКИ
mocks-cli:
	mockery --config .mockery.cli.yml
mocks-domain:
	mockery --config .mockery.domain.yml

# ТЕСТЫ
test-all:
	docker exec -it $(shell docker-compose ps -q app_test) sh -c 'go test ./...'
test-cli:
	docker exec -it $(shell docker-compose ps -q app_test) sh -c 'go test -cover ./internal/cli'
test-domain:
	docker exec -it $(shell docker-compose ps -q app_test) sh -c 'go test -cover ./internal/domain'
test-repository:
	docker exec -it $(shell docker-compose ps -q app_test) sh -c 'go test -cover ./internal/repository'
