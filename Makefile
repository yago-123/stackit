
.PHONY: all
all: lint imports fmt

.PHONY: lint
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

.PHONY: imports
imports:
	@find . -name "*.go" | xargs goimports -w

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -cover ./...

.PHONY: run-redis
run-redis:
	@echo "Running the application..."
	@go run cmd/redis/main.go

.PHONY: run-memory
run-memory:
	@echo "Running the application..."
	@go run cmd/memory/main.go
