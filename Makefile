# ref. http://postd.cc/auto-documented-makefile/

COVERAGE := $(shell mktemp)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run tests only
	go test $$(glide novendor) $(OPT)

test-coverage: ## Run tests and show coverage in browser
	go get github.com/haya14busa/goverage
	goverage -v -coverprofile=$(COVERAGE) -covermode=count
	go tool cover -html=$(COVERAGE)

install: ## Install packages for dependencies
	go get github.com/Masterminds/glide
	glide install
	[ -d .git/hooks ] && cd .git/hooks && [ -L pre-commit ] || ln -s ../../src/scripts/git-hooks/pre-commit
