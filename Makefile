.PHONY: build run test lint clean migrate

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run linter
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf bin/

# Database migrations
migrate:
	go run scripts/migrations/*.go

# Docker commands
docker-build:
	docker build -t yourproject .

docker-run:
	docker-compose up

docker-down:
	docker-compose down

# Generate swagger documentation
swagger:
	swag init -g cmd/api/main.go -o api/swagger