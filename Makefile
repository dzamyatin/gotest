.PHONY: grpc_user

start_user:
	go run ./user/main.go --port 8999

grpc_user:
	protoc \
    --go_out=./user/api --go_opt=paths=source_relative \
    --go-grpc_out=./user/api --go-grpc_opt=paths=source_relative \
    ./user/proto/user.proto
