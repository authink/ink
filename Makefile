.DEFAULT_GOAL := build
ARGS :=

swag:
	go run ./src swag

test:
	APP_ENV=test APP_CWD=$(CURDIR) go test ./src/...
test: swag

build:
	go build -o ./bin/ink ./src
build: test

run:
	APP_ENV=dev ./bin/ink run $(ARGS)
run: build