package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/services"
	"github.com/vuongle/grpc/storages"
	"google.golang.org/grpc"
)

func main() {
	// get port from command line
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	// create laptop server API
	laptopServer := services.NewLaptopServer(storages.NewInMemoryLaptopStore())

	// create grpc server and register laptop server API
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	// start grpc server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server: ", err)
	}
}
