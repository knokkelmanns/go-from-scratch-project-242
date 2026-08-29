install:
	go mod tidy

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size/

test:
	go test ./...

.PHONY: install build test
