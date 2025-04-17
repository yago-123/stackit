
.PHONY: all
all: lint imports fmt run

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

.PHONY: run
run:
	@echo "Running the application..."
	@go run cmd/server/main.go