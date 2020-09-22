#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
PACKAGES_UNITTEST=$(shell go list ./... | grep -v '/simulation' | grep -v '/cli_test')

all: tools install lint

# The below include contains the tools.
include contrib/devtools/Makefile

build: go.sum
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/iris.exe ./cmd/iris
else
	go build $(BUILD_FLAGS) -o build/iris ./cmd/iris
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

build-contract-tests-hooks:
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests.exe ./cmd/contract_tests
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests ./cmd/contract_tests
endif

install: go.sum
	go install $(BUILD_FLAGS) ./cmd/iris

update-swagger-docs: statik
	$(BINDIR)/statik -src=lite/grpc-gateway -dest=lite/grpc-gateway -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
    	echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
.PHONY: update-swagger-docs

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/iris -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

proto-all: proto-tools proto-gen proto-swagger-gen

proto-gen:
	@./scripts/protocgen.sh

proto-swagger-gen:
	@./scripts/protoc-swagger-gen.sh

########################################
### Testing

test: test-unit

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ${PACKAGES_UNITTEST}

lint: golangci-lint
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs gofmt -d -s
	go mod verify

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs goimports -w -local github.com/irisnet/irismod

benchmark:
	@go test -mod=readonly -bench=. ./...
