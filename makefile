DOCKER_COMPOSE=docker compose
GO_APP=go-app
GO_DB=go_db

.PHONY: build
build:
	$(DOCKER_COMPOSE) build

.PHONY: up-db
up-db:
	$(DOCKER_COMPOSE) up -d $(GO_DB)

.PHONY: up-app
up-app:
	$(DOCKER_COMPOSE) up -d $(GO_APP)

.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d

.PHONY: down
down:
	$(DOCKER_COMPOSE) down

.PHONY: rebuild
rebuild:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) build
	$(DOCKER_COMPOSE) up -d

.PHONY: stop
stop:
	$(DOCKER_COMPOSE) stop

.PHONY: restart
restart:
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) up -d

.PHONY: down
down:
	$(DOCKER_COMPOSE) down

.PHONY: up-migrate
up-migrate:
	docker-compose run atlas migrate apply --dir file:///migrations --url "postgres://postgres:postgres@go_db:5432/postgres?sslmode=disable"

# docker-compose run atlas migrate hash
