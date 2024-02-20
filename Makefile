.DEFAULT_GOAL := build
ARGS :=

swag:
	go run ./src swag

test:
	INK_ENV=test INK_CWD=$(CURDIR) go test ./src/...
test: swag

build:
	go build -o ./bin/ink ./src
build: test

run:
	INK_ENV=dev ./bin/ink run $(ARGS)
run: build