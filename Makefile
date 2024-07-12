include .bin-deps.mk
include .env

# СБОРКА
build: .bin-deps # если уже существет network, то проигнорим ошибку
	-docker network create app-network 2>/dev/null || true 
	docker-compose build

# МИГРАЦИЯ
migrate: up-db
	$(GOOSE) -dir db/migrations postgres "$(DATABASE_LOCAL_URL)" up
migrate-test: up-db-test
	$(GOOSE) -dir db/migrations postgres "$(DATABASE_LOCAL_TEST_URL)" up

demigrate: up-db
	$(GOOSE) -dir db/migrations postgres "$(DATABASE_LOCAL_URL)" down
demigrate-test: up-db-test
	$(GOOSE) -dir db/migrations postgres "$(DATABASE_LOCAL_TEST_URL)" down

# ЗАПУСК
up-app:
	docker-compose --env-file .env up -d app
up-db:
	docker-compose --env-file .env up -d db
up-db-test:
	docker-compose --env-file .env up -d db_test
up-broker:
	docker-compose --env-file .env up -d broker
up-broker-test:
	docker-compose --env-file .env up -d broker_test

# ОСТАНОВКА
down:
	docker-compose --env-file .env down
down-app:
	docker-compose --env-file .env down app
down-db:
	docker-compose --env-file .env down db
down-db-test:
	docker-compose --env-file .env down db_test
down-broker:
	docker-compose --env-file .env down broker
down-broker-test:
	docker-compose --env-file .env down broker_test

# ПРОЦЕССЫ
cli: up-app # утилита
	docker exec -it app sh -c './main'

sql: up-db # sql-клиент
	docker exec -it db sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)'
sql-test: up-db-test # sql-клиент для тестов
	docker exec -it db_test sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)_test'

log: up-app # логгер
	-docker exec -it app tail -f ./log.txt

# SHELL
shell-app: up-app
	docker exec -it app sh
shell-db: up-db
	docker exec -it db sh
shell-db-test: up-db-test
	docker exec -it db_test sh
shell-broker: up-broker
	docker exec -it broker sh
shell-broker-test: up-broker-test
	docker exec -it broker_test sh

# МОКИ
mocks:
	$(MOCKERY) --config .mockery.yml

# ТЕСТЫ
tests: up-db-test up-broker-test
	go test ./tests/... -v
tests-unit:
	go test ./tests/unit/... -v
	rm ./tests/unit/cli/log_text.txt
tests-int: up-db-test up-broker-test
	go test ./tests/int/... -v 