.PHONY: run build test clean dev

# Default binary name
BINARY_NAME=oil_selection

build:
	go build -o ${BINARY_NAME} main.go

run:
	ENV=development go run main.go

dev:
	air # For hot-reload during development

test:
	go test -v ./...

clean:
	go clean
	rm -f ${BINARY_NAME}

deps:
	go mod download
	go mod tidy

migrate:
	go run migrations/migrate.go

.DEFAULT_GOAL := run 