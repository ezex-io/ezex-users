PACKAGES=$(shell go list ./... | grep -v 'tests' | grep -v 'grpc/gen')
VERSION=$(shell jq -r 'if .meta != "" then "\(.major).\(.minor).\(.patch)-\(.meta)" else "\(.major).\(.minor).\(.patch)" end' version/version.json)

ifneq (,$(filter $(OS),Windows_NT MINGW64))
EXE = .exe
endif

# Handle sed differences between macOS and Linux
ifeq ($(shell uname),Darwin)
  SED_CMD = sed -i ''
else
  SED_CMD = sed -i
endif

.PHONY: all build clean test proto lint run fmt check proto-check proto-format docker docker-build docker-run

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
	go build -o ./build/main ./internal/cmd/main.go

release:
	go build -ldflags "-s -w" -trimpath -o  ./build/main ./cmd/main.go

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
proto: proto-format
	rm -rf api/gen
	cd api && buf generate --template ./proto/buf.gen.yaml --config ./proto/buf.yaml ./proto

proto-check:
	cd api && buf lint --config ./proto/buf.yaml

proto-format:
	cd api && buf format --config ./proto/buf.yaml -w
########################################
### Run
run: build
	./build/main

########################################
### Docker
docker: docker-build

docker-build:
	@echo "Building Docker image..."
	docker build -t ezex-users:latest .

docker-run:
	@echo "Running Docker container..."
	docker run -d \
		--name ezex-users \
		-p 8080:8080 \
		-p 50051:50051 \
		-e EZEX_USERS_HTTP_SERVER_ADDRESS=":8080" \
		-e EZEX_USERS_GRPC_SERVER_ADDRESS="0.0.0.0:50051" \
		ezex-users:latest
