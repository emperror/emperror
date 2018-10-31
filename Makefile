# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE = github.com/goph/emperror

# Build variables
BUILD_DIR ?= build

DEP_VERSION = 0.5.0
GOLANGCI_VERSION = 1.11.2

.PHONY: setup
setup: vendor ## Setup the project for development

bin/dep: bin/dep-${DEP_VERSION}
	@ln -sf dep-${DEP_VERSION} bin/dep
bin/dep-${DEP_VERSION}:
	@mkdir -p bin
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | INSTALL_DIRECTORY=bin DEP_RELEASE_TAG=v${DEP_VERSION} sh
	@mv bin/dep $@


.PHONY: vendor
vendor: bin/dep ## Install dependencies
	bin/dep ensure -v

.PHONY: clean
clean: ## Clean the working area and the project
	rm -rf bin/ ${BUILD_DIR}/ vendor/

.PHONY: check
check: test lint ## Run tests and linters

.PHONY: test
test: ## Run all tests
	go test ${GOARGS} ./...

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	bin/golangci-lint run

bin/mockery:
	@mkdir -p bin
	GOBIN=${PWD}/bin/ go get github.com/vektra/mockery/cmd/mockery

.PHONY: generate-mocks
generate-mocks: bin/mockery ## Generate test mocks
	bin/mockery -name=Handler -output . -outpkg emperror_test -testonly -case underscore

release-%: ## Release a new version
	@sed -e "s/^## \[Unreleased\]$$/## [Unreleased]\\"$$'\n'"\\"$$'\n'"\\"$$'\n'"## [$*] - $$(date +%Y-%m-%d)/g" CHANGELOG.md > CHANGELOG.md.new
	@mv CHANGELOG.md.new CHANGELOG.md

	@sed -e "s|^\[Unreleased\]: \(.*\)HEAD$$|[Unreleased]: https://${PACKAGE}/compare/v$*...HEAD\\"$$'\n'"[$*]: \1v$*|g" CHANGELOG.md > CHANGELOG.md.new
	@mv CHANGELOG.md.new CHANGELOG.md

ifeq (${TAG}, 1)
	git add CHANGELOG.md
	git commit -s -S -m 'Prepare release v$*'
	git tag -s -m 'Release v$*' v$*
endif

	@echo "Version updated to $*!"
	@echo
	@echo "Review the changes made by this script then execute the following:"
ifneq (${TAG}, 1)
	@echo
	@echo "git add CHANGELOG.md && git commit -S -m 'Prepare release v$*' && git tag -s -m 'Release v$*' v$*"
	@echo
	@echo "Finally, push the changes:"
endif
	@echo
	@echo "git push; git push --tags"

.PHONY: patch
patch: ## Release a new patch version
	@${MAKE} release-$(shell git describe --abbrev=0 --tags | sed 's/^v//' | awk -F'[ .]' '{print $$1"."$$2"."$$3+1}')

.PHONY: minor
minor: ## Release a new minor version
	@${MAKE} release-$(shell git describe --abbrev=0 --tags | sed 's/^v//' | awk -F'[ .]' '{print $$1"."$$2+1".0"}')

.PHONY: major
major: ## Release a new major version
	@${MAKE} release-$(shell git describe --abbrev=0 --tags | sed 's/^v//' | awk -F'[ .]' '{print $$1+1".0.0"}')

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
