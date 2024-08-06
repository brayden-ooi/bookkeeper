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
		go run main.go

start:
		./bin/${BINARY_NAME}

test:
		go test ./...

# make build

miup:
		GOOSE_MIGRATION_DIR="./sql/schema" goose sqlite3 ./db/index.db up

midown:
		GOOSE_MIGRATION_DIR="./sql/schema" goose sqlite3 ./db/index.db down

sqlc:
		sqlc generate