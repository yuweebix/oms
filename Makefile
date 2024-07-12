include .env
ENV_FILE := .env

# СБОРКА
build:
	docker-compose build
build-app:
	docker-compose build app

# NETWORK
create-app-network:
	docker network create app-network

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
up-app:
	docker-compose --env-file $(ENV_FILE) up -d app
up-db:
	docker-compose --env-file $(ENV_FILE) up -d db
up-db-test:
	docker-compose --env-file $(ENV_FILE) up -d db_test
up-broker:
	docker-compose --env-file $(ENV_FILE) up -d broker
up-broker-test:
	docker-compose --env-file $(ENV_FILE) up -d broker_test

# ОСТАНОВКА
down:
	docker-compose --env-file $(ENV_FILE) down
down-app:
	docker-compose --env-file $(ENV_FILE) down app
down-db:
	docker-compose --env-file $(ENV_FILE) down db
down-db-test:
	docker-compose --env-file $(ENV_FILE) down db_test
down-broker:
	docker-compose --env-file $(ENV_FILE) down broker
down-broker-test:
	docker-compose --env-file $(ENV_FILE) down broker_test

# ПРОЦЕССЫ
cli: # утилита
	docker exec -it $(shell docker-compose ps -q app) sh -c './main'

sql: # sql-клиент
	docker exec -it $(shell docker-compose ps -q db) sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)'
sql-test: # sql-клиент для тестов
	docker exec -it $(shell docker-compose ps -q db_test) sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)_test'

log: # логгер
	-docker exec -it $(shell docker-compose ps -q app) tail -f ./log.txt

# SHELL
shell-app:
	docker exec -it $(shell docker-compose ps -q app) sh
shell-db:
	docker exec -it $(shell docker-compose ps -q db) sh
shell-db-test:
	docker exec -it $(shell docker-compose ps -q db_test) sh
shell-broker:
	docker exec -it $(shell docker-compose ps -q broker) sh
shell-broker-test:
	docker exec -it $(shell docker-compose ps -q broker_test) sh

# МОКИ
mocks-cli:
	mockery --config .mockery.cli.yml
mocks-domain:
	mockery --config .mockery.domain.yml

# ТЕСТЫ
test:
	go test ./tests/... -v
test-unit-local:
	go test ./tests/unit/... -v
test-int-local:
	go test ./tests/int/... -v 