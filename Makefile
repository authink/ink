.DEFAULT_GOAL := build
V := 0.1.4
ARGS :=

gen:
	go run ./src gen

tidy:
	go mod tidy
tidy: gen

fmt:
	go fmt ./src/...
fmt: tidy

swag:
	go run ./src swag
swag: fmt

test:
	APP_ENV=test APP_CWD=$(CURDIR) go test ./src/...
test: swag

build:
	go build -o ./bin/ink ./src
build: test

run:
	APP_ENV=dev ./bin/ink run $(ARGS)
run: build

frun:
	APP_ENV=dev ./bin/ink run
frun: tidy

package:
	git tag v$(V)
	git push --tags
package: build