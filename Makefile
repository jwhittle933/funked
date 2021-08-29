pkg ?= strings
build:
	@echo "Building $(pkg)..."
	@go build -o ./bin/$(pkg) ./examples/$(pkg)
.PHONY: build

t ?= ./...
run ?= Test
test:
	@grc &> /dev/null && grc go test -cover $(t) -run=$(run) || go test -cover $(t) -run=$(run)
.PHONY: test

run:
	@./bin/$(pkg)
.PHONY: run