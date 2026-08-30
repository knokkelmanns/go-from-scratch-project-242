install:
	go mod tidy

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size/

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	go test ./...

.PHONY: install build test lint lint-fix test
