BINARY=engine

# Variables
## system
PROJECT_ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
## mockery
MOCKERY_VERSION := $(shell mockery --version 2>/dev/null)
## swag
SWAGGO_VERSION := $(shell swag -v 2>/dev/null)
## lint
GOLANGCILINT_VERSION := $(shell golangci-lint --version 2>/dev/null)
## test cover tmp file
GO_TEST_COVER_TMP_FILE = "/tmp/go-cover.$$.tmp"

## mockery
mockery-install:
ifdef MOCKERY_VERSION
	@echo "Mockery already installed, $(MOCKERY_VERSION)"
else
	go install github.com/vektra/mockery/v2@latest
endif
	@exit 0

mockery:
ifndef MOCKERY_VERSION
	@echo "Mockery not installed. Please run \"make mockery-install\""
else
	rm -rf "$(PROJECT_ROOT_DIR)/entities/mocks"
	cd "$(PROJECT_ROOT_DIR)/entities" && mockery --all --keeptree
endif
	@exit 0

## swaggo
swaggo-install:
ifdef SWAGGO_VERSION
	@echo "Swaggo already installed, $(SWAGGO_VERSION)"
else
	go install github.com/swaggo/swag/cmd/swag@latest
endif
	@exit 0

swaggo:
ifndef SWAGGO_VERSION
	@echo "Swaggo not installed. Please run \"make swaggo-install\""
else
	cd "$(PROJECT_ROOT_DIR)" && swag init -g ./app/main.go && swag fmt -g ./app/main.go
endif
	@exit 0

## lint
lint-install:
ifdef GOLANGCILINT_VERSION
	@echo "golangci-lint already installed, $(GOLANGCILINT_VERSION)"
else
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
endif
	@exit 0

lint:
ifndef GOLANGCILINT_VERSION
	@echo "Swaggo not installed. Please run \"make lint-install\""
else
	cd "$(PROJECT_ROOT_DIR)" && golangci-lint run -v ./...
endif
	@exit 0

## test
test: 
	cd "$(PROJECT_ROOT_DIR)" && go test -v -cover -covermode=atomic ./...
	@exit 0

test-cover:
	cd "$(PROJECT_ROOT_DIR)" && go test -coverprofile=$(GO_TEST_COVER_TMP_FILE) ./... && go tool cover -html=$(GO_TEST_COVER_TMP_FILE) && unlink $(GO_TEST_COVER_TMP_FILE)
	@exit 0

# Script
install:
	$(shell cd "$(PROJECT_ROOT_DIR)" && go get ./app)
	$(shell go mod tidy)
ifndef MOCKERY_VERSION
	make mockery-install
endif
ifndef SWAGGO_VERSION
	make swaggo-install
endif
ifndef GOLANGCILINT_VERSION
	make lint-install
endif

ifneq ($(shell swag -v 2>/dev/null), $(shell golangci-lint --version 2>/dev/null), $(shell mockery --version 2>/dev/null), )
	@echo "Ready to go."
endif

	@exit 0

before-run:
	make mockery
	make swaggo
	make lint
	make test
	@exit 0

run:
	cd "$(PROJECT_ROOT_DIR)" && go run ./app