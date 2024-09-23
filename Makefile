.PHONY: help build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## build docker image to deploy
	docker build -t ysaito8015/go_mysql_example:${DOCKER_TAG} \
		--target deploy ./


build-local: ## build docker image to local development
	docker compose build --no-cache

up: ## start docker compose
	docker compose up -d

down: ## stop docker compose
	docker compose down

logs: ## show logs
	docker compose logs -f

ps: ## show docker compose status
	docker compose ps

test: ## run test
	go test -race -shuffle=on ./...

dry-migrate: ## Try migration
	mysqldef -u mysql -p password -h 127.0.0.1 -P 33306 example -dry-run < ./_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u mysql -p password -h 127.0.0.1 -P 33306 example < ./_tools/mysql/schema.sql

help: ## show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

