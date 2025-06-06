BINARY_NAME = ezex-users
BUILD_DIR = build
CMD_DIR = ./internal/cmd/server/

# Default target
all: build test

########################################
### Tools needed for development
devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

proto:
	@mkdir -p ./pkg/grpc
	protoc --go_out=./pkg/grpc --go_opt paths=source_relative \
		   --go-grpc_out=./pkg/grpc --go-grpc_opt paths=source_relative \
		   --proto_path=./pkg/proto ./pkg/proto/*.proto

# SQLC generate sql
sqlc:
	sqlc generate .

mock:
	mockgen -source=./internal/port/database.go	-destination=./internal/mock/mock_database.go	-package=mock

########################################
### Building
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

release:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-s -w" -trimpath -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(BUILD_DIR)

########################################
### Testing
# Run only unit tests
test:
	go test ./... -short

# Run only integration tests
test-integration:
	go test ./internal/test/...

########################################
### Formatting the code
fmt:
	@echo "Formatting code..."
	@go tool gofumpt -l -w .

lint:
	@echo "Running lint..."
	@go tool golangci-lint  run ./... --timeout=20m0s

check: fmt lint

########################################
### Run
run: build
	./build/$(BINARY_NAME)

########################################
### Docker
docker:
	docker build --tag ezex-users .

docker-build:
	@echo "Building Docker image..."
	docker build -t ezex-users:latest .

docker-run:
	@echo "Running Docker container..."
	docker run -d \
		--name ezex-users \
		ezex-users:latest

.PHONY: docker docker-build docker-run mock sqlc
.PHONY: devtools proto docker
.PHONY: test test-integration
.PHONY: fmt lint check
.PHONY: run build release clean
