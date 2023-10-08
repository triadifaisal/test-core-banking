DATABASE_SOURCE_NAME='postgres://user_banking:secret_banking@localhost:5433/banking_test?sslmode=disable'

always-run:

vendor: always-run
	env GO111MODULE=on go mod vendor

build-base-image:
	@docker build build/docker/base/ --no-cache -t test/corebanking-base:0.0.1-corebanking

migrate:
	# https://github.com/golang-migrate/migrate
	migrate -path database/migrations -database $(DATABASE_SOURCE_NAME) up

migrate-create:
	# https://github.com/golang-migrate/migrate
	migrate create -ext sql -dir database/migrations -seq $(DATABASE_SOURCE_NAME)

rollback:
	# https://github.com/golang-migrate/migrate
	migrate -path database/migrations -database $(DATABASE_SOURCE_NAME) down 1
