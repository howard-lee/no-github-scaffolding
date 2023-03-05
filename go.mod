module blocktype

go 1.18

require (
	api v1.0.0
	server v1.0.0
)

require (
	github.com/go-kit/kit v0.12.0
	github.com/stretchr/testify v1.8.1
	google.golang.org/grpc v1.53.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	grpcData v1.0.0 // indirect
)

replace api v1.0.0 => ./api/v1

replace server v1.0.0 => ./internal/grpc/server

replace test v1.0.0 => ./test

replace grpc_test v1.0.0 => ./test/grpc_test

replace grpcData v1.0.0 => ./internal/grpc/grpcData

