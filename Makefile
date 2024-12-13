.PHONY: help build tests fmt format-lines lint

include .env

GO_LINES_IGNORED_DIRS=
GO_PACKAGES=./main/...
GO_FOLDERS=$(shell echo ${GO_PACKAGES} | sed -e "s/\.\///g" | sed -e "s/\/\.\.\.//g")

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Compile the binary
	@mkdir -p bin
	@go build -buildmode=plugin -o lib/$(APP_NAME).so $(GO_PACKAGES)

tests: ## runs all tests
	go test ./... -covermode=atomic

fmt: ## formats all go files
	go fmt ./...
	make format-lines

format-lines: ## formats all go files with golines
	go install github.com/segmentio/golines@latest
	golines -w -m 120 --ignore-generated --shorten-comments --ignored-dirs=${GO_LINES_IGNORED_DIRS} ${GO_FOLDERS}

lint: ## runs all linters
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...
