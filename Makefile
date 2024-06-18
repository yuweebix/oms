include .env
ENV_FILE := .env

# Build the project
build:
	docker-compose build

# Database commands
up-db:
	docker-compose --env-file $(ENV_FILE) up -d db

down-db:
	docker-compose --env-file $(ENV_FILE) down db

start-db:
	docker-compose --env-file $(ENV_FILE) up db

stop-db:
	docker-compose --env-file $(ENV_FILE) down db

# Application commands
up-app:
	docker-compose --env-file $(ENV_FILE) up -d app

down-app:
	docker-compose --env-file $(ENV_FILE) down app

start-app:
	docker-compose --env-file $(ENV_FILE) up app

stop-app:
	docker-compose --env-file $(ENV_FILE) down app

# All services commands
up:
	docker-compose --env-file $(ENV_FILE) up -d

down:
	docker-compose --env-file $(ENV_FILE) down

start:
	docker-compose --env-file $(ENV_FILE) up

stop:
	docker-compose --env-file $(ENV_FILE) down

# Open a shell in the running app container
shell-app:
	docker exec -it $$(docker-compose ps -q app) sh

# Open a shell in the running db container
shell-db:
	docker exec -it $$(docker-compose ps -q db) sh

# Run database migrations
migrate:
	docker-compose --env-file $(ENV_FILE) run --rm app sh -c 'goose -dir db/migrations postgres "$(DATABASE_URL)" up'
