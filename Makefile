.PHONY: all test get-deps

all:
	go install

test: 
	go test ./...

get-deps:
	go get -u github.com/satori/go.uuid
	go get -u github.com/stretchr/testify
