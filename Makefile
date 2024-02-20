test:
	INK_ENV=test INK_CWD=$(CURDIR) go test ./src/...

build:
	go build -o ./bin/ink ./src

build: test