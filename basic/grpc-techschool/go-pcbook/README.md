# Prerequisites for gRPC in Go

Official and latest doc and example

```
https://grpc.io/docs/languages/go/quickstart/

https://github.com/grpc/grpc-go/tree/master/examples
```

## Install protoc compiler

```
https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/
```

## Install the protocol compiler plugins for Go

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Generate go code from proto file

```
protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative proto/*.proto
```

## Viewing until

https://www.youtube.com/watch?v=i9H3BaRGLEc&list=PLy_6D98if3UJd5hxWNfAqKMr15HZqFnqf&index=15

## Some usaful commands

Find which process running on port 8080

```
netstat -ano|findstr "PID :8080"
```

Kill process

```
taskkill /PID 18264 /f
```
