DOCKER_COMPOSE_KAFKA=docker-compose -f docker-compose-kafka.yml -p core-banking
DATABASE_SOURCE_NAME='postgres://user_banking:secret_banking@localhost:5433/banking_test?sslmode=disable'

always-run:

vendor: always-run
	env GO111MODULE=on go mod vendor

build-base-image:
	@docker build build/docker/base/ --no-cache -t test/corebanking-base:0.0.1-corebanking

create-network:
	@docker network create -d bridge corebanking-network

migrate:
	# https://github.com/golang-migrate/migrate
	migrate -path database/migrations -database $(DATABASE_SOURCE_NAME) up

migrate-force:
	# https://github.com/golang-migrate/migrate
	migrate -path database/migrations -database $(DATABASE_SOURCE_NAME) force $(version)

migrate-create:
	# https://github.com/golang-migrate/migrate
	migrate create -ext sql -dir database/migrations -seq $(name)

rollback:
	# https://github.com/golang-migrate/migrate
	migrate -path database/migrations -database $(DATABASE_SOURCE_NAME) down 1

start-dev:
	@echo "Start dev..."
	@echo
	@$(DOCKER_COMPOSE_KAFKA) up -d
	@echo "Wait for 5 seconds for database up and running properly"
	@sleep 5
	@docker compose up

stop-dev:
	@echo "Stop dev..."
	@echo
	@$(DOCKER_COMPOSE_KAFKA) stop
	@$(DOCKER_COMPOSE) stop

start-dev-consumer:
	@docker-compose -p test-core-banking run --name=consumer_1 --rm console go run ./cmd/console/consumer/main.go -topic MutationIDs