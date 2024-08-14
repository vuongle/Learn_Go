gen:
	protoc --grpc-gateway_out=pb --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

clean:
	del .\pb\*.go

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server:
	go run cmd/server/main.go -port 8080

rest:
	go run cmd/server/main.go -port 8081 -type rest

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover ./...

.PHONY: clean gen server client test