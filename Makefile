pkg ?= strings
build:
	@echo "Building $(pkg)..."
	@go build -o ./bin/$(pkg) ./examples/$(pkg)
.PHONY: build

test:
	@grc go test -cover ./$(pkg)
.PHONY: test

run:
	@./bin/$(pkg)
.PHONY: run