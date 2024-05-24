package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	// thiết lập kết nối với gRPC service
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
