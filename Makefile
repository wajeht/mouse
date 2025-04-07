push:
	@make format
	@make test
	@git auto

test:
	@go test ./...

build:
	@go build -v ./...

run:
	@go run ./cmd/main.go

format:
	@go fmt ./...

tidy:
	@go mod tidy

deps:
	@go get -u ./...

clean:
	@go clean

