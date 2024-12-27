.PHONY: build test

build:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/ncmdl ./ncmdl

test:
	go test ./...
