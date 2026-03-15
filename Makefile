BINARY=loglint
CUSTOM_GCL=custom-gcl
PACKAGE?=./...

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ./cmd/loglint $(PACKAGE)

.PHONY: custom-lint
custom-lint:
	golangci-lint custom

.PHONY: plugin
plugin: custom-lint
	./$(CUSTOM_GCL) run $(PACKAGE)

.PHONY: build
build:
	go build -o ./bin/$(BINARY) ./cmd/loglint

.PHONY: run-binary
run-binary: build
	./bin/$(BINARY) $(PACKAGE)

.PHONY: clean
clean:
	rm -rf ./bin ./$(CUSTOM_GCL)