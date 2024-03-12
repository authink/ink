.DEFAULT_GOAL := build
V := 0.1.8
ARGS :=

tidy:
	go mod tidy

gen:
	go run ./src gen -m github.com/authink/ink/src/orm/models
gen: tidy

fmt:
	go fmt ./src/...
fmt: gen

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
frun: gen

package:
	git tag v$(V)
	git push --tags
package: build