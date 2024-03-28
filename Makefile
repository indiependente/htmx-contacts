deps:
	go mod tidy

build:
	go build -o bin/$(shell basename $(PWD)) .

run:
	go run .

dev:
	air
