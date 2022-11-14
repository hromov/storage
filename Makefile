BINARY_NAME=storage

build:
	go build -o ${BINARY_NAME} cmd/main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

lint:
	golangci-lint run

generate:
	go generate ./...

start:
	go run cmd/main.go