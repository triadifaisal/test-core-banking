###########################################################################
## Make file
## @author: Faisal Triadi <triadifaisal@gmail.com>
## @since: 2023.10.06
###########################################################################

DOCKER_COMPOSE=docker-compose -f docker-compose.yml -p corebanking

build-base-image:
	@docker build build/docker/base/ --no-cache -t test/corebanking-base:0.0.1-corebanking

init-dev:
	rm -f config.ini
	rm -f .env
	@cp cicd/config/local/config.ini config.ini
	@cp cicd/config/local/.env .env

start-dev:
	@echo "############################"
	@echo "####### Core Banking #######"
	@echo "############################"
	@echo
	@echo "Start dev..."
	@echo
	@make init-dev
	# @docker build -f "build/docker/base/Dockerfile" -t "test/corebanking-base:0.0.1-corebanking" .
	# @$(DOCKER_COMPOSE_DEV) up -d database redis
	# @$(DOCKER_COMPOSE_DEV) up -d elasticsearch
	# @echo "Wait for 20 seconds for database and es up and running properly"
	# @sleep 20
	@$(DOCKER_COMPOSE) up

swaggo-testing:
	rm -rf internal/auth/handler/rest/docs
	swag init -o internal/auth/handler/rest/docs --parseDependency --parseInternal -d internal/auth/handler/rest -g gin_server.go