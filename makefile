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

ENV_FILE=cfg/.env

define build_db_url
"postgres://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=$${DB_SSLMODE}"
endef

migrate-up:
	@set -o allexport; source cfg/.env; \
	migrate -path ./internal/migrations -database $(call build_db_url) up

migrate-down:
	@set -o allexport; source cfg/.env; \
	migrate -path ./internal/migrations -database $(call build_db_url) down

migrate-version:
	@set -o allexport; source cfg/.env; \
	migrate -path ./internal/migrations -database $(call build_db_url) version

migrate-force:
	@set -o allexport; source cfg/.env; \
	migrate -path ./internal/migrations -database $(call build_db_url) force

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./internal/migrations -seq $$name

# === Help ===

help:
	@echo "🛠️  Cappit Makefile Commands:"
	@echo ""
	@echo "🧱  Dev Environment:"
	@echo "  make up            - Start cappit, Postgres, and Redis containers"
	@echo "  make up-dev        - Start only Redis and Postgres (for local dev use)"
	@echo "  make down          - Stop and remove all containers"
	@echo "  make restart       - Restart all containers"
	@echo "  make logs          - Tail logs from containers"
	@echo "  make ps            - Show container status"
	@echo ""
	@echo "🧪  Development:"
	@echo "  make test          - Run Go tests in isolated container"
	@echo "  make lint          - Run linter using golangci-lint container"
	@echo ""
	@echo "📦  Database Migrations:"
	@echo "  make migrate-up     - Apply all up migrations"
	@echo "  make migrate-down   - Revert the last migration"
	@echo "  make migrate-version - Show current migration version"
	@echo "  make migrate-force  - Force set a specific migration version"
	@echo "  make migrate-create - Create a new migration file"

