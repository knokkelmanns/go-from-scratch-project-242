install:
	go mod tidy
	cd tests && go mod tidy

build:
	mkdir -p bin
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size/

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	cd tests && go test ./...

.PHONY: install build test lint lint-fix
