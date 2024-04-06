BINARY_NAME=bookkeeper
.DEFAULT_GOAL := build

fmt:
		go fmt cmd/root.go
.PHONY:fmt

lint: fmt
		golangci-lint run
.PHONY:lint

impt: lint
		goimports -w cmd/root.go
.PHONY:impt

vet: impt
		go vet
.PHONY:vet
		# shadow cmd/root.go

build: vet
		go build -o bin/${BINARY_NAME}
.PHONY:build

clean:
		go clean

run:
		./bin/${BINARY_NAME}

test:
		go test ./...

# make build