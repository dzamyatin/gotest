.PHONY: grpc_user, migrate_create, start_console, start_user, tst

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

test:
	go test -v ./user/...
tst:
	go test -v ./tst/...
test_b:
	go test -bench=MyFunction -benchtime=10s -benchmem
#main test
#unit test
#benchmark test
#integration ?
#exaple test
