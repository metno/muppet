.PHONY: all test

all:
	go install

test: 
	go test ./...
