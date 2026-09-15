run:
	go run ./src/cmd/hello

build:
	go build -o bin/hello ./src/cmd/hello

test:
	go test ./...