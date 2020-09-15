test:
	go test -cover ./...

run:
	go run cmd/main.go

build:
	mkdir bin
	go build cmd/main.go
	chmod +x main
	mv main bin/task
	tar -zcvf task.tar.gz bin