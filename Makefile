ifneq ($(wildcard docker/.env.example),)
	ENV_FILE = .env.example
endif
ifneq ($(wildcard .env.example),)
	ifeq ($(COMPOSE_PROJECT_NAME),)
		include .env.example
	endif
endif
ifneq ($(wildcard docker/.env),)
	ENV_FILE = .env
endif
ifneq ($(wildcard .env),)
	ifeq ($(COMPOSE_PROJECT_NAME),)
		include .env
	endif
endif

export

.SILENT:
.PHONY: help
help: ## Display this help screen
	awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

.PHONY: install
install: ## Installations
	go mod download
	go mod verify

.PHONY: lint
lint: ## Run linters
	golangci-lint run

.PHONY: run
run: ## Run application
	go run .

.PHONY: compose-convert
compose-convert: ## Converts the compose file to platform's canonical format
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) convert

.PHONY: compose-build
compose-build: ## Build or rebuild services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) build

.PHONY: compose-up
compose-up: ## Create and start containers
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) up -d

.PHONY: compose-down
compose-down: ## Stop and remove containers, networks
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) down --remove-orphans

.PHONY: compose-logs
compose-logs: ## View output from containers
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) logs -f

.PHONY: compose-ps
compose-ps: ## List containers
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) ps

.PHONY: compose-ls
compose-ls: ## List running compose projects
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) ls

.PHONY: compose-exec
compose-exec: ## Execute a command in a running container
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) exec backend bash

.PHONY: compose-start
compose-start: ## Start services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) start

.PHONY: compose-restart
compose-restart: ## Restart services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) restart

.PHONY: compose-stop
compose-stop: ## Stop services
	docker-compose -f docker/docker-compose.yml --env-file docker/$(ENV_FILE) stop

.PHONY: docker-rm-volume
docker-rm-volume: ## Remove db volume
	docker volume rm -f fiber_clean_db_data

.PHONY: docker-clean
docker-clean: ## Remove unused data
	docker system prune
