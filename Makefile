.PHONY: grpc_user, migrate_create, start_console, start_user

start_user:
	go run ./user/main.go --port 8999

start_console:
	go run ./user/console/main.go $(ARG)

schema_update:
	go run ./user/cmd/main.go gorm_schema_migration

migrate_create:
	go run ./user/cmd/main.go migrate_create

migrate:
	go run ./user/cmd/main.go migrate

grpc_user:
	protoc \
    --go_out=./user/api --go_opt=paths=source_relative \
    --go-grpc_out=./user/api --go-grpc_opt=paths=source_relative \
    ./user/proto/user.proto
