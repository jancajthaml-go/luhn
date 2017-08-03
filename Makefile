.PHONY: all
all: build test benchmark

.PHONY: build
build:
	@go build luhn.go

.PHONY: test
test:
	@go test

.PHONY: benchmark
benchmark:
	@go test -bench=.