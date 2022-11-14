BINARY_NAME=storage

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
 	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
 	GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

test:
	go test ./...

lint:
	golangci-lint run

generate:
	go generate ./...