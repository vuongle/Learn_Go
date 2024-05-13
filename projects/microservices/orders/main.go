package main

import (
	"context"
	common "go-microservices/commons"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	// Create gRPC server
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	// Create service and store
	store := NewStore()
	svc := NewService(store)

	NewGRPCHandler(grpcServer)

	svc.CreateOrder(context.Background())

	// Start gRPC server
	log.Printf("gRPC Server running at %s", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
