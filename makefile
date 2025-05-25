.PHONY: lint
lint:  ## Run the Go linter
	docker-compose run --rm linter

.PHONY: test
test:  ## Run Go tests
	docker-compose run --rm test

.PHONY: build
build:  ## Build the chatbot Docker image
	docker-compose build chatbot

.PHONY: up
up:  ## Start the chatbot service
	docker-compose up -d chatbot

.PHONY: down
down:  ## Stop and remove all services
	docker-compose down

.PHONY: help
help:  ## Show this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)