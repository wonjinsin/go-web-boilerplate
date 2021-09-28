PACKAGE = pikachu
CUSTOM_OS = ${GOOS}
BASE_PATH = $(shell pwd)
BIN = $(BASE_PATH)/bin
BINARY_NAME = bin/$(PACKAGE)
MAIN = $(BASE_PATH)/main.go
GOLINT = $(BIN)/golint
MOCK = $(BIN)/mockgen
PKG_LIST = $(shell cd $(BASE_PATH) && cat pkg.list)

ifneq (, $(CUSTOM_OS))
	OS ?= $(CUSTOM_OS)
else
	OS ?= $(shell uname | awk '{print tolower($0)}')
endif
build:
	GOOS=$(OS) go build -o $(BINARY_NAME) $(MAIN)

.PHONY: vet
vet:
	go vet

.PHONY: fmt
fmt:
	go fmt

.PHONY: lint
lint: build-lint
	$Q $(GOLINT) $(PKG_LIST)

build-lint:
	go list ./... > pkg.list
	GOBIN=$(BIN) go get golang.org/x/lint/golint
	go mod vendor

.PHONY: test
test:
	go test -v -cover ./...

test-all: vet fmt lint

build-gomod:
	[ -f ./go.mod ] || go mod init $(PACKAGE)

build-mocks:
	GOBIN=$(BIN) go get github.com/golang/mock/mockgen
	$(MOCK) -source=service/service.go -destination=mock/mock_service.go -package=mock
	$(MOCK) -source=repository/repository.go -destination=mock/mock_repository.go -package=mock

.PHONY: vendor
vendor: build-gomod \
	build-mocks
	go mod vendor

start:
	@$(BIN)/$(PACKAGE)

all: vendor build

clear:; $(info cleaning…) @ 
	@rm -rf vendor mock bin
	@rm go.mod go.sum pkg.list