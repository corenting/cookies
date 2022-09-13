VERSION=$(shell git rev-parse --short HEAD)

format:
	golangci-lint run --fix ./...

style:
	golangci-lint run ./...

run:
	go run main.go

build:
	CGO_ENABLED=0 go build -ldflags "-s -X 'github.com/corenting/cookies/internal/entities.AppVersion=$(VERSION)'" -o bin/debug/cookies

test:
	go test ./...

coverage:
	go test -coverprofile coverage.out ./...
