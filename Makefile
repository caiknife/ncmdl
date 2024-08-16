.PHONY: build test

build:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./dist/ncmdl

test:
	go test ./...
