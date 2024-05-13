package main

import (
	common "go-microservices/commons"
	"log"
	"net/http"

	pb "go-microservices/commons/api"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr          = common.EnvString("HTTP_ADDR", ":8080")
	ordersServiceAddr = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fail to dial orders service: %v", err)
	}
	defer conn.Close()

	log.Printf("Dialing orders service at %s", ordersServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()

	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Gateway server running on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Fail to started the gateway server")
	}
}
