# Variables
DOCKER_COMPOSE=docker compose
GO_APP=go-app
GO_DB=go_db

# Build the Go app and database services
.PHONY: build
build:
	$(DOCKER_COMPOSE) build

# Run the database service
.PHONY: up-db
up-db:
	$(DOCKER_COMPOSE) up -d $(GO_DB)

# Run the Go app service
.PHONY: up-app
up-app:
	$(DOCKER_COMPOSE) up -d $(GO_APP)

# Run both services independently (not dependent on each other)
.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d

# Stop and remove containers, networks, volumes for both services
.PHONY: down
down:
	$(DOCKER_COMPOSE) down

# Rebuild the Go app and database services
.PHONY: rebuild
rebuild:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) build
	$(DOCKER_COMPOSE) up -d

# Stop the containers without removing them
.PHONY: stop
stop:
	$(DOCKER_COMPOSE) stop

# Restart containers independently
.PHONY: restart
restart:
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) up -d

# Show logs for all services
.PHONY: logs
logs:
	$(DOCKER_COMPOSE) logs -f

# Show logs for Go app service
.PHONY: logs-app
logs-app:
	$(DOCKER_COMPOSE) logs -f $(GO_APP)

# Show logs for database service
.PHONY: logs-db
logs-db:
	$(DOCKER_COMPOSE) logs -f $(GO_DB)
