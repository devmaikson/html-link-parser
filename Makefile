BINARY_NAME=html-link-parser

build:
	mkdir -p bin
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=window go build -o bin/${BINARY_NAME}-windows main.go

run:
	bin/${BINARY_NAME}

build_and_run: build run

test:
	go test ./...

clean:
	go clean
	rm bin/${BINARY_NAME}-darwin
	rm bin/${BINARY_NAME}-linux
	rm bin/${BINARY_NAME}-windows
