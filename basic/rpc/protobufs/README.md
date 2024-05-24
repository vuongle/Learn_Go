### Tạo RPC code sử dụng hàm dựng sẵn trong gRPC plugin từ protoc-gen-go

protoc --go_out=. --go_opt=paths=source_relative hello.proto

### Tạo gRPC code sử dụng hàm dựng sẵn trong gRPC plugin từ protoc-gen-go

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hello.proto
