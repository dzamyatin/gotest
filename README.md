
user 
- api -> generated grpc files
- handler -> grpc handlers to implement grpc logic
- proto -> contain grpc protobuf files
- entity -> contain project entities
- repository -> implement logic to create entities, to work with database
- use_case -> actions


# Prometheus
https://prometheus.io/docs/guides/go-application/



go mod init github.com/dzamyatin/gotest
go mod tidy
go mod vendor





go build -gcflags=-m main.go

#Trace
https://habr.com/ru/articles/742402/


# Mock
mockgen -source=/home/dzamyatin/GolandProjects/newGame/vendor/gorm.io/gorm/gorm.go -destination=user/internal/mock_gorm.go -package=app


#Table test t.Run(...)
https://habr.com/ru/companies/otus/articles/741618/


To confirm if autovacuum is enabled, you can check the postgresql.conf file in your PostgreSQL installation directory (e.g., /etc/postgresql/12/main/postgresql.conf).