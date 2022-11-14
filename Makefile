BINARY_NAME=storage

build:
	go build -o ${BINARY_NAME} main.go

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

lint:
	golangci-lint run

generate:
	go generate ./...