# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Dev variables
GO_SOURCE_FILES = $(shell find . -type f -name "*.go" -not -name "bindata.go" -not -path "./vendor/*" -not -path "./mocks/*")
GO_PACKAGES = $(shell go list ./... | grep -v /vendor | grep -v /mocks)

.PHONY: setup
setup:: dep ## Setup the project for development

.PHONY: dep
dep: ## Install dependencies
	@glide install

.PHONY: clean
clean:: ## Clean the working area
	rm -rf vendor/

.PHONY: check
check:: test cs ## Run tests and linters

PASS=$(shell printf "\033[32mPASS\033[0m")
FAIL=$(shell printf "\033[31mFAIL\033[0m")
COLORIZE=sed ''/PASS/s//${PASS}/'' | sed ''/FAIL/s//${FAIL}/''

.PHONY: test
test: ## Run unit tests
	@go test -tags '${TAGS}' ${ARGS} ${GO_PACKAGES} | ${COLORIZE}

.PHONY: watch-test
watch-test: ## Watch for file changes and run tests
	reflex -t 2s -d none -r '\.go$$' -- $(MAKE) ARGS="${ARGS}" test

.PHONY: cs
cs: ## Check that all source files follow the Go coding style
	@gofmt -l ${GO_SOURCE_FILES} | read something && echo "Code differs from gofmt's style" 1>&2 && exit 1 || true

.PHONY: csfix
csfix: ## Fix Go coding style violations
	@gofmt -l -w -s ${GO_SOURCE_FILES}

.PHONY: envcheck
envcheck:: ## Check environment for all the necessary requirements
	$(call executable_check,Go,go)
	$(call executable_check,Glide,glide)
	$(call executable_check,Reflex,reflex)

define executable_check
    @printf "\033[36m%-30s\033[0m %s\n" "$(1)" `if which $(2) > /dev/null 2>&1; then echo "\033[0;32m✓\033[0m"; else echo "\033[0;31m✗\033[0m"; fi`
endef

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
