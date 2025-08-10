# Redis
up:
	docker compose up -d

down:
	docker compose down

restart: down up

# golang app
run:
	@go run cmd/main.go