BINARY_NAME = ezex-users
BUILD_DIR = build
CMD_DIR = internal/cmd/server/main.go

# Default target
all: build test

########################################
### Tools needed for development
devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

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
test:
	go test ./...

########################################
### Formatting the code
fmt:
	gofumpt -l -w .

check:
	golangci-lint run --timeout=20m0s

.PHONY: devtools
.PHONY: build release
.PHONY: test
.PHONY: fmt check
	go build -o ./bin/ezex-users$(EXE) ./cmd/server.go

build_race:
	go build -race -o ./bin/ezex-users$(EXE) ./cmd/server.go
########################################
### Proto
proto:
	protoc --go_out ./api/grpc/proto --go_opt paths=source_relative \
           --go-grpc_out ./api/grpc/proto --go-grpc_opt paths=source_relative \
	       --proto_path=./api/grpc/proto api/grpc/proto/*.proto
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

.PHONY: docker docker-build docker-run
.PHONY: devtools proto docker
.PHONY: build release
.PHONY: test
.PHONY: fmt check