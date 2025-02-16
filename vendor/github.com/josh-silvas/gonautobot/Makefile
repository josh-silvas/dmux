APP_NAME  := gonautobot
REPO      := github.com/josh-silvas

COMPOSE   := docker-compose -p $(APP_NAME) --project-directory "develop"  -f "develop/docker-compose.yml"
RUN       := $(COMPOSE) run --rm develop

VERSION   := v$(shell cat VERSION)
HASH      := $(shell git rev-parse HEAD)
TS        := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

default: help

cli: .env ## Launch a local development/testing environment.
	@$(RUN) bash
.PHONY: cli

##################################################################################################################
# LINTING AND TESTING
##################################################################################################################
tests: .env ## Run all code quality, unit and integration testing.
	@$(RUN) make _lint
	@$(RUN) make _unittest
	@echo "✅ Testing completed! ✅"
.PHONY: test


cover: ## Run the unit-test coverage report.
	@go tool cover -html=coverage.txt
.PHONY: cover

lint: .env ## Run golang-ci lint on all sub-packages
	@$(RUN) make _lint
.PHONY: lint

_lint:
	@echo "🔧 Running golangci-lint... 🔧"
	@golangci-lint run --tests=false --exclude-use-default=false
.PHONY: _lint

unittest: .env ## Run UnitTest only.
	@$(RUN) make _unittest
.PHONY: unittest

_unittest:
	@echo "⏳ Running Golang Unit Tests... ⏳"
	@go test -v -short -coverprofile=coverage.txt -covermode=atomic  $(REPO)/$(APP_NAME)/... -tags=unit | { grep -v 'no test files'; true; }
.PHONY: _unittest

##################################################################################################################
# EXTRAS
##################################################################################################################
.env:
	@if [ ! -f "develop/.env" ]; then \
	   echo "Creating environment file...\nPLEASE OVERRIDE VARIABLES IN develop/.env WITH YOUR OWN VALUES!"; \
	   cp develop/example.env develop/.env; \
	fi
.PHONY: .env

help:
	@echo "\033[1m\033[01;32m\
	$(shell echo $(APP_NAME) | tr  '[:lower:]' '[:upper:]') $(VERSION) \
	\033[00m\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' \
	$(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help
