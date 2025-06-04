.PHONY: help

SHELL := /bin/bash
VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
NOW=$(shell date +"%Y-%m-%d_%H:%M:%S")
COMMIT_REF=$(shell git rev-parse --short HEAD)
BUILD_ARGS=-ldflags "-X github.com/timo-reymann/SchemaNest/pkg/buildinfo.GitSha=$(COMMIT_REF) -X github.com/timo-reymann/SchemaNest/pkg/buildinfo.Version=$(VERSION) -X github.com/timo-reymann/SchemaNest/pkg/buildinfo.BuildTime=$(NOW)"
BIN_PREFIX="dist/"
BIN_PREFIX_SCHEMA_REGISTRY="$(BIN_PREFIX)schema-nest-registry-"
BIN_PREFIX_SCHEMA_UPLOADER="$(BIN_PREFIX)schema-nest-cli-"
CMD_REGISTRY = "./cmd/schema-nest-registry"
CMD_UPLOADER = "./cmd/schema-nest-cli"

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

build-linux: create-dist ## Build binaries for linux
	@GOOS=linux GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-amd64 $(BUILD_ARGS) $(CMD_REGISTRY)
	@GOOS=linux GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-i386 $(BUILD_ARGS) $(CMD_REGISTRY)
	@GOOS=linux GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-arm $(BUILD_ARGS) $(CMD_REGISTRY)
	@GOOS=linux GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)linux-arm64 $(BUILD_ARGS) $(CMD_REGISTRY)

	@GOOS=linux GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)linux-amd64 $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=linux GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)linux-i386 $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=linux GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)linux-arm $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=linux GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)linux-arm64 $(BUILD_ARGS) $(CMD_UPLOADER)

build-windows: create-dist ## Build binaries for windows
	@GOOS=windows GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)windows-amd64.exe $(BUILD_ARGS) $(CMD_REGISTRY)
	@GOOS=windows GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)windows-i386.exe $(BUILD_ARGS) $(CMD_REGISTRY)
	@GOOS=windows GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)windows-arm.exe $(BUILD_ARGS) $(CMD_REGISTRY)

	@GOOS=windows GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)windows-amd64.exe $(BUILD_ARGS) $(CMD_UPLOADER)
	@GOOS=windows GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)windows-i386.exe $(BUILD_ARGS) $(CMD_UPLOADER)
	@GOOS=windows GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)windows-arm.exe $(BUILD_ARGS) $(CMD_UPLOADER)

build-darwin: create-dist  ## Build binaries for macOS
	@GOOS=darwin GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)darwin-amd64 $(BUILD_ARGS) $(CMD_REGISTRY)
	@GOOS=darwin GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)darwin-arm64 $(BUILD_ARGS) $(CMD_REGISTRY)

	@GOOS=darwin GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)darwin-amd64 $(BUILD_ARGS) $(CMD_UPLOADER)
	@GOOS=darwin GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)darwin-arm64 $(BUILD_ARGS) $(CMD_UPLOADER)

build-freebsd: create-dist ## Build binaries for FreeBSD
	@GOOS=freebsd GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)freebsd-amd64 $(BUILD_ARGS) $(CMD_REGISTRY)
    @GOOS=freebsd GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)freebsd-i386 $(BUILD_ARGS) $(CMD_REGISTRY)
    @GOOS=freebsd GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)freebsd-arm64 $(BUILD_ARGS) $(CMD_REGISTRY)
    @GOOS=freebsd GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)freebsd-arm $(BUILD_ARGS) $(CMD_REGISTRY)

	@GOOS=freebsd GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)freebsd-amd64 $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=freebsd GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)freebsd-i386 $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=freebsd GOARCH=arm64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)freebsd-arm64 $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=freebsd GOARCH=arm go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)freebsd-arm $(BUILD_ARGS) $(CMD_UPLOADER)

build-openbsd: create-dist ## Build binaries for OpenBSD
	@GOOS=openbsd GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)openbsd-amd64 $(BUILD_ARGS) $(CMD_REGISTRY)
    @GOOS=openbsd GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_REGISTRY)openbsd-i386 $(BUILD_ARGS) $(CMD_REGISTRY)

    @GOOS=openbsd GOARCH=amd64 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)openbsd-amd64 $(BUILD_ARGS) $(CMD_UPLOADER)
    @GOOS=openbsd GOARCH=386 go build -o $(BIN_PREFIX_SCHEMA_UPLOADER)openbsd-i386 $(BUILD_ARGS) $(CMD_UPLOADER)

create-checksums: ## Create checksums for binaries
	@find ./dist -type f -exec sh -c 'sha256sum {} | cut -d " " -f 1 > {}.sha256' {} \;

go-generate: ## Generate go files for migrations and openapi code
	@go generate pkg/internal/ecosystems/generate.go

build: go-generate build-linux build-darwin build-windows build-freebsd build-openbsd create-checksums ## Build binaries for all platform
