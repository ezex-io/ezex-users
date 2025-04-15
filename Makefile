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
