.PHONY: build
build:
	go build -v ./cmd/bigfile/main

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_goal := build