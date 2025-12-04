# Makefile for Go projects
# Usage: make <target>
# Common targets: help build run test lint fmt vet clean docker

# --- Configurable vars ---
MODULE       := $(shell basename $(CURDIR))
BINARY       := server
MAIN         := cmd/server/main.go
BUILD_DIR    := dist
BIN_DIR      := bin
LDFLAGS      := -s -w
TIME         := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GIT_TAG      := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
GOOS         ?= $(shell go env GOOS)
GOARCH       ?= $(shell go env GOARCH)
CGO_ENABLED  ?= 0

# Docker image name (optional)
DOCKER_IMAGE := $(MODULE):latest

# Tools (used by `install-tools`)
GOLANGCI_LINT_VERSION := v1.59.0
SWAG_VERSION          := v1.8.12

# --- Derived ---
OUTPUT := $(BIN_DIR)/$(BINARY)

# --- Phony ---
.PHONY: help fmt vet lint test build build-release build-linux build-multi run deps \
	install-tools generate swag docker docker-build docker-push clean

# --- Help ---
help:
	@echo "Makefile targets:"
	@echo "  make deps              # download go module dependencies"
	@echo "  make install-tools     # install local dev tools (golangci-lint, swag)"
	@echo "  make fmt               # gofmt / goimports (use gofmt by default)"
	@echo "  make vet               # go vet"
	@echo "  make lint              # run golangci-lint"
	@echo "  make test              # run unit tests"
	@echo "  make generate          # run go:generate"
	@echo "  make swag              # generate swagger docs (swag)"
	@echo "  make build             # build binary for host OS/ARCH"
	@echo "  make build-release     # build binary with version ldflags"
	@echo "  make build-linux       # build linux/amd64 static binary"
	@echo "  make build-multi       # build linux/amd64 and linux/arm64"
	@echo "  make run               # run from source"
	@echo "  make docker            # build docker image"
	@echo "  make docker-push       # push docker image (set DOCKER_REGISTRY env)"
	@echo "  make clean             # clean build artifacts"

# --- Dependency management ---
deps:
	@echo "Downloading module dependencies..."
	go mod download

# --- Developer tools installer ---
install-tools:
	@echo "Installing dev tools..."
	# golangci-lint
	@echo " - golangci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	# swag (optional)
	@echo " - swag"
	@go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

# --- Format / Vet / Lint ---
fmt:
	@echo "Formatting code..."
	go fmt ./...

vet:
	@echo "Running go vet..."
	go vet ./...

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...

# --- Tests ---
test:
	@echo "Running tests..."
	go test ./... -v

# --- Code generation / docs ---
generate:
	@echo "Running go generate..."
	go generate ./...

swag:
	@echo "Generating Swagger docs (swag)..."
	swag init -g $(MAIN) -o docs --parseDependency --parseInternal || true

# --- Build ---
build: deps
	@mkdir -p $(BIN_DIR)
	@echo "Building for $(GOOS)/$(GOARCH)..."
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT) $(MAIN)

build-release: deps
	@mkdir -p $(BUILD_DIR)
	@echo "Building release binary with version info..."
	@LD_FLAGS="$(LDFLAGS) -X 'main.Version=$(GIT_TAG)' -X 'main.Commit=$(GIT_COMMIT)' -X 'main.BuildTime=$(TIME)'" ; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$$LD_FLAGS" -o $(BUILD_DIR)/$(BINARY)-linux-amd64 $(MAIN)

build-linux: deps
	@mkdir -p $(BUILD_DIR)
	@echo "Building linux/amd64 static binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-linux-amd64 $(MAIN)

build-multi: deps
	@mkdir -p $(BUILD_DIR)
	@echo "Cross-building linux/amd64 and linux/arm64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-linux-amd64 $(MAIN)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64  go build -o $(BUILD_DIR)/$(BINARY)-linux-arm64  $(MAIN)

# --- Run ---
run:
	@echo "Running $(MAIN)..."
	go run $(MAIN)

# --- Docker ---
docker: build-release
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .

docker-build:
	@echo "Building Docker image without building binary first..."
	docker build -t $(DOCKER_IMAGE) .

docker-push:
ifndef DOCKER_REGISTRY
	$(error DOCKER_REGISTRY is not set. Example: make docker-push DOCKER_REGISTRY=myrepo)
endif
	@IMAGE=$(DOCKER_REGISTRY)/$(DOCKER_IMAGE)
	@docker tag $(DOCKER_IMAGE) $$IMAGE
	@docker push $$IMAGE
	@echo "Pushed $$IMAGE"

# --- Clean ---
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR) $(BUILD_DIR) docs/swagger

