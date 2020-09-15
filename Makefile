test:
	go test -cover ./...

run:
	go run cmd/main.go

build:
	go build cmd/main.go