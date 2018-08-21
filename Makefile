# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

DEP_VERSION = 0.5.0
GOLANGCI_VERSION = 1.9.3

.PHONY: setup
setup:: dep ## Setup the project for development

bin/dep: ## Install dep
	@mkdir -p ./bin/
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | INSTALL_DIRECTORY=./bin DEP_RELEASE_TAG=v${DEP_VERSION} sh

.PHONY: dep
dep: bin/dep ## Install dependencies
	@bin/dep ensure

.PHONY: clean
clean:: ## Clean the working area
	rm -rf bin/ build/ vendor/

.PHONY: check
check:: test lint ## Run tests and linters

PASS=$(shell printf "\033[32mPASS\033[0m")
FAIL=$(shell printf "\033[31mFAIL\033[0m")
COLORIZE=sed ''/PASS/s//${PASS}/'' | sed ''/FAIL/s//${FAIL}/''

.PHONY: test
test: ## Run unit tests
	@go test -tags '${TAGS}' ${ARGS} ./... | ${COLORIZE}

bin/golangci-lint: ## Install golangci linter
	@mkdir -p ./bin/
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	@bin/golangci-lint run

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
