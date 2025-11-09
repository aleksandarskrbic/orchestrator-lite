.PHONY: build run test clean fmt vet

build:
	go build -o bin/orchestrator-lite main.go

run:
	go run main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test
