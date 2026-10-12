.PHONY: build test fmt check clean

build:
	mkdir -p bin
	go build -o bin/codekb ./cmd/codekb

test:
	go test ./...

fmt:
	gofmt -w cmd internal

check:
	test -z "$$(gofmt -l cmd internal)"
	go vet ./...
	go test ./...
	go build ./cmd/codekb

clean:
	go clean

