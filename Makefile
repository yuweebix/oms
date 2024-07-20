include .bin-deps.mk
include .vendor-proto.mk
include .env

# ЗАВИСИМОСТИ
deps: .bin-deps .bin-protoc-deps

# СБОРКА
build:
# если уже существет network, то проигнорим ошибку
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
up-redis:
	docker-compose --env-file .env up -d redis
up-redis-test:
	docker-compose --env-file .env up -d redis_test
up-metrics:
	docker-compose --env-file .env up grafana

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
down-redis:
	docker-compose --env-file .env down redis
down-redis-test:
	docker-compose --env-file .env down redis_test

# ПРОЦЕССЫ
server: up-app # сервер
	docker exec -it app sh -c './main'

sql: up-db # sql-клиент
	docker exec -it db sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)'
sql-test: up-db-test # sql-клиент для тестов
	docker exec -it db_test sh -c 'psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)_test'

redis-cli: up-redis
	docker exec -it redis redis-cli
redis-cli-test: up-redis-test
	docker exec -it redis_test redis-cli

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
shell-redis: up-redis
	docker exec -it redis sh
shell-redis-test: up-redis-test
	docker exec -it redis_test sh

# МОКИ
mocks:
	$(MOCKERY) --config .mockery.yml

# ТЕСТЫ
tests: up-db-test up-broker-test up-redis-test
	go test ./tests/... -v
tests-unit:
	go test ./tests/unit/... -v
tests-int: up-db-test up-broker-test up-redis-test
	go test ./tests/int/... -v 

# gRPC
generate: .vendor-proto generate-no-deps

generate-no-deps:
	mkdir -p gen/orders/v1/proto
	mkdir -p gen/orders/v1/swagger
	mkdir -p gen/returns/v1/proto
	mkdir -p gen/returns/v1/swagger
	$(PROTOC) -I api/proto/orders/v1 \
		-I vendor.proto \
		api/proto/orders/v1/orders.proto \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) --go_out=./gen/orders/v1/proto --go_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) --go-grpc_out=./gen/orders/v1/proto --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) --grpc-gateway_out=./gen/orders/v1/proto --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPI) --openapiv2_out=./gen/orders/v1/swagger \
		--plugin=protoc-gen-validate=$(PROTOC_GEN_VALIDATE) --validate_out=lang=go,paths=source_relative:./gen/orders/v1/proto \
		--experimental_allow_proto3_optional=true
	$(PROTOC) -I api/proto/returns/v1 \
		-I vendor.proto \
		api/proto/returns/v1/returns.proto \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) --go_out=./gen/returns/v1/proto --go_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) --go-grpc_out=./gen/returns/v1/proto --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) --grpc-gateway_out=./gen/returns/v1/proto --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPI) --openapiv2_out=./gen/returns/v1/swagger \
		--plugin=protoc-gen-validate=$(PROTOC_GEN_VALIDATE) --validate_out=lang=go,paths=source_relative:./gen/returns/v1/proto \
		--experimental_allow_proto3_optional=true