test:
	go test -cover ./...

run:
	go run cmd/main.go

build:
	mkdir dist
	go build cmd/main.go
	chmod +x main
	mv main dist/task
	tar -zcvf task.tar.gz dist