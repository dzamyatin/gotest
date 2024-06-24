.PHONY: grpc_user, migrate_create, start_console, start_user, tst, grpc, trace, trace_read

grpc:
	GOGC=70 GOMEMLIMIT=50MiB go run ./user/main.go --port 8999

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

#curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1
#https://golangci-lint.run/
lint:
	golangci-lint run ./...
	go vet ./...
	shadow ./...

#benchmark test https://habr.com/ru/articles/268585/
test_b:
	go test -race -bench=. -benchtime=1s -benchmem -cpuprofile=cpu.out -memprofile=mem.out -v ./tst/search
#go tool pprof -svg ./perftest00.test ./cpu.out > cpu.svg
test_b_see:
	go tool pprof -http=:8088 cpu.out
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
tst:
	go test -v ./tst/...
#go test ./... -coverprofile=coverage.txt
tst_coverage:
	go test -v ./tst/... -coverprofile="coverage.out"
#go tool cover -html coverage.txt -o index.html
#https://habr.com/ru/companies/badoo/articles/301990/
tst_coverage_res:
	go tool cover -html="coverage.out"


prof_see:
	go tool pprof  -http=:8088 ./user/var/cpuProfiler.prof
prof_mem_see:
	go tool pprof  -http=:8088 ./user/var/memProfiler.prof


trace_read:
	go tool trace -http "0.0.0.0:8088" ./tracetest trace.out
trace:
	curl -o trace.out -u test:test http://localhost:8998/debug/pprof/trace?seconds=10



