.PHONY: dev up down wait-db build run test

DB_HOST=localhost
DB_PORT=5432
DB_USER=pg

dev: up wait-db
	air

up:
	docker-compose up -d database

down:
	docker-compose down

wait-db:
	@echo "Waiting for Postgres to be ready..."
	@until docker-compose exec -T database pg_isready -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT); do \
		echo "Postgres is unavailable - sleeping"; \
		sleep 2; \
	done
	@echo "Postgres is up!"

build: test
	@echo "Building..."
	go build -o ./api ./cmd/api

run: up wait-db build
	@echo "Running..."
	ENVIRONMENT=production go run ./...

test:
	@echo "Running tests..."
	go test ./tests/... -v
