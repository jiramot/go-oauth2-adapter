BINARY=build

build:
	go build -o bin/engine ./cmd/main.go

.PHONY: build