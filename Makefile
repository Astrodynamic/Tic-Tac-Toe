.DEFAULT_GOAL := help

GO ?= go

CMD_DIR := ./cmd/server
BIN_DIR := ./bin
BIN_NAME := server
BIN := $(BIN_DIR)/$(BIN_NAME)

.PHONY: help
help: ## Show available targets
	@awk 'BEGIN {FS=":.*## "; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_.-]+:.*## / {printf "  %-14s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run
run: ## Run the server (go run)
	@$(GO) run $(CMD_DIR)

.PHONY: build
build: ## Build the server binary into ./bin/
	@mkdir -p "$(BIN_DIR)"
	@$(GO) build -trimpath -o "$(BIN)" $(CMD_DIR)

.PHONY: test
test: ## Run all tests
	@$(GO) test ./...

.PHONY: fmt
fmt: ## Format code (gofmt + gofmt on module packages)
	@$(GO) fmt ./...
	@gofmt -w .

.PHONY: tidy
tidy: ## Sync go.mod/go.sum
	@$(GO) mod tidy

.PHONY: clean
clean: ## Remove build artifacts
	@rm -rf "$(BIN_DIR)"
