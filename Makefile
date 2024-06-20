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
tst_coverage:
	go test -v ./tst/... -coverprofile="coverage.out"
tst_coverage_res:
	go tool cover -html="coverage.out"

#benchmark test https://habr.com/ru/articles/268585/
test_b:
	go test -bench=. -benchtime=1s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -v ./tst/search
#go test -bench=. -benchmem bench_test.go > new.txt
#git stash
#go test -bench=. -benchmem bench_test.go > old.txt
#go get golang.org/x/tools/cmd/benchcmp
#benchcmp old.txt new.txt
# Cpu profile:
# go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out bench_test.go


#main test
#unit test
#integration ?
#exaple test

#go test ./... -cover
#go test ./... -coverprofile=coverage.txt
#go tool cover -html coverage.txt -o index.html
