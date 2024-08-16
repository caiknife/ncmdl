.PHONY: build test

build:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter

test:
	go test ./...
