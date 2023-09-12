.DEFAULT_GOAL:=help

.PHONY: build
build: clean ## Build binaries
	@go build -o jsondup ./cmd/jsondup/main.go

.PHONY: clean
clean: ## Clean up binaries
	@rm -f jsondup

.PHONY: cover
cover: test ## Display test coverage report
	@go tool cover -func=coverage.txt

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: ## Lint source code
	@golangci-lint run ./...

.PHONY: test
test: ## Run unit tests
	@go test -v -coverprofile=coverage.txt ./...
