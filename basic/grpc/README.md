## Steps to create gRPC server

1. Create .proto file: Define message and service
2. Generate gRPC go code by using "protoc" compiler
3. Install missing packages

```
go get -u google.golang.org/grpc
go mod tidy
```
