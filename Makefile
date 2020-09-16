test:
	go test -cover ./...

run:
	go run cmd/main.go

build:
	mkdir dist
	GOOS=darwin GOARCH=amd64 go build cmd/main.go
	mv main dist/task-mac
	GOOS=linux GOARCH=amd64 go build cmd/main.go
	mv main dist/task-linux
	chmod +x dist/task-mac
	chmod +x dist/task-linux
	tar -zcvf task.tar.gz dist