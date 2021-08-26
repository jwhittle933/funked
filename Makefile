pkg ?= strings
build:
	@echo "Building $(pkg)..."
	@go build -o ./bin/$(pkg) ./examples/$(pkg)
.PHONY: build

t ?= ./...
test:
	@grc &> /dev/null && grc go test -cover $(t) || go test -cover $(t)
.PHONY: test

run:
	@./bin/$(pkg)
.PHONY: run