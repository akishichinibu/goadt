GOLANGCI_VERSION := v2.1.5

init/golangci-lint:
	@echo "Installing golangci-lint..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)

tidy:
	@go mod tidy
	@go mod vendor

build: tidy
	@go build

test: tidy gen
	@go test ./...

lint: tidy
	@$(shell go env GOPATH)/bin/golangci-lint run ./... --timeout 5m

gen: tidy
	@go generate ./...
