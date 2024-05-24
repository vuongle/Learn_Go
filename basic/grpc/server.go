package main

import (
	"context"
	"demo-grpc/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Create a struct that implements the "InvoicerServer" interface in invoicer_grpc.pb.go
type InvoicerServerImpl struct {
	invoicer.UnimplementedInvoicerServer
}

func (is InvoicerServerImpl) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
			Pdf:  []byte("pdf"),
			Docx: []byte("docx"),
		},
		nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("can not create tcp listener %v", err)
	}

	grpcServer := grpc.NewServer()
	service := &InvoicerServerImpl{}
	invoicer.RegisterInvoicerServer(grpcServer, service)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("can not serve %v", err)
	}
}
