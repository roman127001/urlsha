SHELL:=bash

default: help

.PHONY: help
help: ## Available commands
	@clear
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[0;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@echo ""


##@ Local

.PHONY: run
run: ## Run the application
	cd cmd && go run *.go

.PHONY: test
test: ## Run the tests
	go test -v ./...



##@ Docker

.PHONY: run-docker
run-docker: ## Run the application in a docker container
	docker-compose up -d

.PHONY: stop-docker
stop-docker: ## Stop the application in a docker container
	docker-compose down
