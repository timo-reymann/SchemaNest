.PHONY: help

SHELL := /bin/bash
VERSION=$(shell git describe --tags)
NOW=$(shell date +"%Y-%m-%d_%H:%M:%S")
COMMIT_REF=$(shell git rev-parse --short HEAD)
BUILD_ARGS=-ldflags "-X github.com/timo-reymann/SchemaNest/pkg/buildinfo.GitSha=$(COMMIT_REF) -X github.com/timo-reymann/SchemaNest/pkg/buildinfo.Version=$(VERSION) -X github.com/timo-reymann/SchemaNest/pkg/buildinfo.BuildTime=$(NOW)" -tags prod
BIN_PREFIX="dist/"
BIN_PREFIX_SCHEMA_REGISTRY="$(BIN_PREFIX)schema-nest-registry-"
BIN_PREFIX_SCHEMA_CLI="$(BIN_PREFIX)schema-nest-cli-"
CMD_REGISTRY = "./cmd/schema-nest-registry"
CMD_CLI = "./cmd/schema-nest-cli"

clean: ## Cleanup artifacts
	@rm -rf dist/

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

coverage: ## Run tests and measure coverage
	@go test -covermode=count -coverprofile=/tmp/count.out -v ./...

test-coverage-report: coverage ## Run test and display coverage report in browser
	@go tool cover -html=/tmp/count.out

save-coverage-report: coverage ## Save coverage report to coverage.html
	@go tool cover -html=/tmp/count.out -o coverage.html

create-dist: ## Create dist folder if not already existent
	@mkdir -p dist/

build-ui: ## Build UI
	@cd ui && yarn build

build-linux: create-dist ## Build binaries for linux
	@CGO_ENABLED=1 CC="zig cc -target x86_64-linux-musl" CXX="zig c++ -target x86_64-linux-musl"  GOOS=linux GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-amd64 $(BUILD_ARGS) $(CMD_REGISTRY)
	@CGO_ENABLED=1 CC="zig cc -target x86-linux-musl" CXX="zig c++ -target x86-linux-musl"  GOOS=linux GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-i386 $(BUILD_ARGS) $(CMD_REGISTRY)
	@CGO_ENABLED=1 CC="zig cc -target arm-linux-musleabihf" CXX="zig c++ -target arm-linux-musleabihf"  GOOS=linux GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-arm $(BUILD_ARGS) $(CMD_REGISTRY)

	@CGO_ENABLED=1 CC="zig cc -target x86_64-linux-musl" CXX="zig c++ -target x86_64-linux-musl"  GOOS=linux GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_CLI)linux-amd64 $(BUILD_ARGS) $(CMD_CLI)
	@CGO_ENABLED=1 CC="zig cc -target arm-linux-musleabihf" CXX="zig c++ -target arm-linux-musleabihf"  GOOS=linux GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_CLI)linux-arm $(BUILD_ARGS) $(CMD_CLI)
    @CGO_ENABLED=1 CC="zig cc -target x86-linux-musl" CXX="zig c++ -target x86-linux-musl"  GOOS=linux GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_CLI)linux-i386 $(BUILD_ARGS) $(CMD_CLI)

build-windows: create-dist ## Build binaries for windows
	@CGO_ENABLED=1 CC="zig cc -target x86_64-windows-gnu" CXX="zig c++ -target x86_64-windows-gnu"  GOOS=windows GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)windows-amd64.exe $(BUILD_ARGS) $(CMD_REGISTRY)
	@CGO_ENABLED=1 CC="zig cc -target x86_64-windows-gnu" CXX="zig c++ -target x86_64-windows-gnu"  GOOS=windows GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_CLI)windows-amd64.exe $(BUILD_ARGS) $(CMD_CLI)

build-darwin: create-dist  ## Build binaries for macOS
	@CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)darwin-amd64 $(BUILD_ARGS) $(CMD_REGISTRY)
	@CGO_ENABLED=1 GOOS=darwin go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)darwin-arm64 $(BUILD_ARGS) $(CMD_REGISTRY)

	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_CLI)darwin-amd64 $(BUILD_ARGS) $(CMD_CLI)
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_CLI)darwin-arm64 $(BUILD_ARGS) $(CMD_CLI)

create-checksums: ## Create checksums for binaries
	@find ./dist -type f -exec sh -c 'sha256sum {} | cut -d " " -f 1 > {}.sha256' {} \;

go-generate: ## Generate go files for migrations and openapi code
	@go generate pkg/internal/ecosystems/generate.go

build-image-cli: ## Build the CLI container image
	@docker buildx build . \
		-t timoreymann/schemanest-cli:${VERSION} \
		-t timoreymann/schemanest-cli:latest \
		-f docker/cli.Dockerfile \
		--platform linux/amd64,linux/arm64 \
		--push

build-image-registry: ## Build the registry container image
	@docker buildx build . \
		-t timoreymann/schemanest-registry:${VERSION} \
		-t timoreymann/schemanest-registry:latest \
		-f docker/registry.Dockerfile \
		--platform linux/amd64,linux/arm64 \
		--push

build-image: build-image-cli build-image-registry ## Build all images

build: go-generate build-linux build-darwin build-windows create-checksums ## Build binaries for all platform
