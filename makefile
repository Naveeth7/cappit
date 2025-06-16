up:
	@echo "🔼 Starting all containers (cappit, postgres, redis)..."
	docker-compose up -d postgres redis cappit

up-dev:
	@echo "🔼 Starting only postgres and redis..."
	docker-compose up -d postgres redis

down:
	@echo "🔽 Stopping all containers..."
	docker-compose down

restart:
	@echo "♻️  Restarting all containers..."
	docker-compose down
	docker-compose up -d postgres redis cappit

logs:
	docker-compose logs -f

ps:
	docker-compose ps

# === Test & Lint ===

test:
	@echo "🧪 Running tests in container..."
	docker-compose run --rm test

lint:
	@echo "🔍 Running linter in container..."
	docker-compose run --rm linter

# === Help ===

help:
	@echo "🛠️  Cappit Makefile Commands:"
	@echo "  make up         - Start cappit, postgres, redis containers"
	@echo "  make up-dev     - Start only redis and postgres (for local dev use)"
	@echo "  make down       - Stop and remove all containers"
	@echo "  make restart    - Restart all containers"
	@echo "  make logs       - Tail logs from containers"
	@echo "  make ps         - Show container status"
	@echo "  make test       - Run Go tests in isolated container"
	@echo "  make lint       - Run linter using golangci-lint container"
