REPO ?= github.com/josh-silvas
PROJECT_NAME ?= dmux
VERSION := $(shell cat VERSION)
COMPOSE := docker-compose --project-name $(PROJECT_NAME) --project-directory "development"  -f "development/docker-compose.yml"
RUNNER := $(COMPOSE) run --rm $(PROJECT_NAME)-build
UNAME := $(shell uname -m)

default: help

cli: .env
	@$(RUNNER) bash
.PHONY: cli

ci: .env ## Build the CI environment.
	@$(COMPOSE) build --no-cache
.PHONY: ci

# -------------------------------------------------------------------------------------------
# RELEASE: Goreleaser release functionality
# -------------------------------------------------------------------------------------------
release: .env ## Release a new version and build, tag.
	@make _release
.PHONY: release

_release:
	@rm -rf build/*
	@git tag -d $(VERSION) || true
	@git tag $(VERSION)
	@goreleaser --rm-dist
.PHONY: _release

test_release: .env ## Test a release
	@make _test_release
.PHONY: test_release

_test_release:
	@git tag -d $(VERSION) || true
	@git tag $(VERSION)
	@goreleaser --rm-dist --snapshot
	@git tag -d $(VERSION) || true
.PHONY: _test_release

# -------------------------------------------------------------------------------------------
# CODE-QUALITY/TESTS: Linting and testing directives.
# -------------------------------------------------------------------------------------------
tests: .env ## Run all code quality, unit and integration testing.
	@$(RUNNER) make _lint
	@$(RUNNER) make _unittest
.PHONY: test

lint: .env ## Run golint on all sub-packages
	@$(RUNNER) make _lint
.PHONY: lint

_lint:
	@echo "Running YAML..."
	@yamllint -c .yamllint.yml .
	@echo "Running golangci-lint..."
	@golangci-lint run --tests=false --exclude-use-default=false
.PHONY: _lint

unittest: .env ## Run UnitTest only.
	@$(RUNNER) make _unittest
.PHONY: unittest

_unittest:
	@go test -v -short -coverprofile=coverage.txt -covermode=atomic ./...  | { grep -v 'no test files'; true; }
.PHONY: _unittest

# -------------------------------------------------------------------------------------------
# DOCUMENTATION: Doc builders and processes.
# -------------------------------------------------------------------------------------------
docs: .env ## Build mkdocs documentation
	@$(RUNNER) make _docs
.PHONY: docs

_docs:
	@mkdocs build
	@mkdocs gh-deploy --force --clean
.PHONY: _docs

# -------------------------------------------------------------------------------------------
# DOCKER: Building of containers and pushing to registries
# -------------------------------------------------------------------------------------------
image: .env ## Builds a go binary and docker container with goreleaser.
	@$(COMPOSE) build --no-cache
.PHONY: image

pull: _env-REGISTRY _env-REPO _env-PROJECT_NAME ## Pulls container from GCR.
	@docker pull ${REPO}/${REGISTRY}/${PROJECT_NAME}-build:latest
.PHONY: pull

push: _env-REGISTRY _env-REPO _env-PROJECT_NAME ## Pushes the docker container to GCR.
	@docker push ${REPO}/${REGISTRY}/${PROJECT_NAME}-build:latest
.PHONY: push

# -------------------------------------------------------------------------------------------
# HELPERS: Helper declarations
# -------------------------------------------------------------------------------------------
.env:
	@if [ ! -f "development/.env" ]; then \
	   echo "Creating environment file...\nPLEASE OVERRIDE VARIABLES IN development/.env WITH YOUR OWN VALUES!"; \
	   cp development/example.env development/.env; \
	fi
.PHONY: .env

_env-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		echo "Please check README.md or Makefile for variables required."; \
		echo "(╯°□°）╯︵ ┻━┻"; \
		exit 1; \
	fi
.PHONY: _env-%

help:
	@echo "\033[1m\033[01;32m\
	$(shell echo $(PROJECT_NAME) | tr  '[:lower:]' '[:upper:]') $(VERSION) \
	\033[00m\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' \
	$(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help
