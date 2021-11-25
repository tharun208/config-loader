GO_TEST := go test $(GOFLAGS) $(LD_FLAGS)
GO_TEST_OPTS ?=
PKG_LIST := ./...
GOPATH_DIR := $(shell go env GOPATH | awk -F: '{print $$1}')
GOPATH_BIN_DIR := $(GOPATH_DIR)/bin
export PATH := $(GOPATH_BIN_DIR):$(PATH)

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z0-9_/-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dev/install/goimports
dev/install/goimports: ## Dev: Install goimports
	go get golang.org/x/tools/cmd/goimports

.PHONY: fmt
fmt: ## Dev: Run go fmt
	go fmt $(GOFLAGS) ./...

.PHONY: vet
vet: ## Dev: Run go vet
	go vet $(GOFLAGS) ./...

.PHONY: imports
imports: ## Dev: Runs goimports in order to organize imports
	goimports -w -local config-loader -d `find . -type f -name '*.go' -not -path './vendored/*'`

.PHONY: check
check: fmt vet imports ## Dev: Run code checks(go fmt, go vet, ...)
	git diff --quiet || test $$(git diff --name-only | grep -v -e 'go.mod$$' -e 'go.sum$$' | wc -l) -eq 0 || ( echo "The following changes (result of code checks) have been detected:" && git --no-pager diff && false ) # fail if Git working tree is dirty

.PHONY: test/unit
test/unit: ## Dev: Run unit tests
	$(GO_TEST) $(GO_TEST_OPTS) -covermode=atomic -coverpkg=./... $(PKG_LIST)

.PHONY: build/app
build/app: ## Dev: Build cmd app
	go build -o config-loader ./app/cmd


.PHONY: run
run: ## Dev: Run cmd app
	go run ./app/cmd/main.go

