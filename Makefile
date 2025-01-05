VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT_SHA ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS = -X 'yourproject/pkg/version.Version=${VERSION}' \
          -X 'yourproject/pkg/version.CommitSHA=${COMMIT_SHA}' \
          -X 'yourproject/pkg/version.BuildTime=${BUILD_TIME}'

.PHONY: build run test lint clean migrate release

# Build the application
build:
	go build -ldflags "${LDFLAGS}" -o bin/api cmd/api/main.go

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

release:
	GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o bin/api-linux-amd64 cmd/api/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o bin/api-darwin-amd64 cmd/api/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o bin/api-windows-amd64.exe cmd/api/main.go
