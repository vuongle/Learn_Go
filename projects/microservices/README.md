#### Step 1: Init mod in each service

```
go mod init go-microservices/[service name]
```

#### Step 2: Create a go module-workspace

```
go work init .\common\ .\gateway\ .\kitchen\ .\orders\ .\payments\ .\stock\
```

#### Step 3: Running a service from module-workspace

At the root of "microservices" folder, run:

```
go run ./gateway
```

Where, "gateway": name of a mod

#### Used packages

In "gateway" folder

```
go get github.com/joho/godotenv
go install github.com/cosmtrek/air@latest
air init
air
```

#### gRPC

https://grpc.io/docs/languages/go/quickstart/

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
