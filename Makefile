APP_NAME=otp-auth
DOCKER_COMPOSE=docker compose
MIGRATE_CMD=$(DOCKER_COMPOSE) run --rm migrate

.PHONY: build run logs stop migrate-new migrate-up migrate-down psql redis

## Build the Go binary (inside Docker)
build:
	$(DOCKER_COMPOSE) build app

## Run the full stack (app + postgres + redis)
run:
	$(DOCKER_COMPOSE) up --build

## Stop all containers
stop:
	$(DOCKER_COMPOSE) down

## View app logs
logs:
	$(DOCKER_COMPOSE) logs -f app

## Create a new migration (example: make migrate-new name=add_users)
migrate-new:
	@docker run --rm -v $(PWD)/migrations:/migrations migrate/migrate \
		create -ext sql -dir /migrations -seq $(name)

## Apply migrations
migrate-up:
	$(MIGRATE_CMD) up

## Roll back last migration
migrate-down:
	$(MIGRATE_CMD) down 1

## Open psql shell
psql:
	$(DOCKER_COMPOSE) exec postgres psql -U otpuser -d otp_auth

## Open redis-cli shell
redis:
	$(DOCKER_COMPOSE) exec redis redis-cli
