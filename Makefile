all:  build run

run:
	bin/main

build:
	go build cmd/server/main.go 


lint:
	golangci-lint run ./...

test:
	go test ./...

