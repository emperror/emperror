# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

DEP_VERSION = 0.5.0
GOLANGCI_VERSION = 1.10.2

.PHONY: setup
setup: vendor ## Setup the project for development

bin/dep: bin/dep-${DEP_VERSION}
bin/dep-${DEP_VERSION}:
	@mkdir -p bin
	@rm -rf bin/dep-*
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | INSTALL_DIRECTORY=./bin DEP_RELEASE_TAG=v${DEP_VERSION} sh
	@touch $@

.PHONY: vendor
vendor: bin/dep ## Install dependencies
	@bin/dep ensure

.PHONY: clean
clean: ## Clean the working area
	rm -rf bin/ build/ vendor/

.PHONY: check
check: test lint ## Run tests and linters

.PHONY: test
test: ## Run tests
	go test ${ARGS} ./...

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	@rm -rf bin/golangci-lint-*
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@touch $@

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	@bin/golangci-lint run

bin/mockery:
	@mkdir -p bin
	GOBIN=${PWD}/bin/ go get github.com/vektra/mockery/cmd/mockery

.PHONY: generate-mocks
generate-mocks: bin/mockery ## Generate test mocks
	bin/mockery -name=Handler -output . -outpkg emperror_test -testonly -case underscore

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
