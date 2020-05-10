.PHONY: build
build:
	go build ./...

.PHONY: tests
tests:
	go test ./... -v
	
	
	