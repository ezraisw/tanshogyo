.PHONY: protoc
protoc:
	protoc --proto_path=./api/proto/tanshogyo --go_out=./ --go-grpc_out=./ ./api/proto/tanshogyo/*.proto
.PHONY: dideps
dideps:
	go run github.com/google/wire/cmd/wire show ./internal/app
.PHONY: di
di:
	go run github.com/google/wire/cmd/wire ./internal/app
.PHONY: setup
setup:
	go mod vendor
	go run github.com/google/wire/cmd/wire ./internal/app
.PHONY: build
build:
	go build -o ./bin/app ./cmd/app
.PHONY: test
test:
	go test ./...
