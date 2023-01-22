test:
	go test -v ./...

build: test
	go build cmd/main.go -o bin/generator