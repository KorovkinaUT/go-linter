BINARY=./bin/loglint
CUSTOM_GCL=./custom-gcl
PACKAGE ?= ./...

.PHONY: build
build: build-cli build-plugin

.PHONY: build-cli
build-cli:
	@go build -o $(BINARY) ./cmd/loglint

.PHONY: build-plugin
build-plugin:
	@golangci-lint custom

.PHONY: run
run:
	@$(BINARY) $(PACKAGE)

.PHONY: plugin
plugin:
	@$(CUSTOM_GCL) run $(PACKAGE)

.PHONY: test
test:
	@go test ./...

.PHONY: clean
clean:
	@rm -rf ./bin $(CUSTOM_GCL)