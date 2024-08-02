package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vuongle/grpc/client"
	interceptors_client "github.com/vuongle/grpc/client/interceptors/client"
	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 60 * time.Second
)

func authMethods() map[string]bool {
	const laptopServicePath = "/grpcservice.laptop.LaptopService/"
	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

// / This is the gRPC client
func main() {
	// get server address from command line
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dialing server: %s", *serverAddress)

	// create grpc connection
	conn1, err := grpc.NewClient(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// create auth client
	authClient := client.NewAuthClient(conn1, username, password)
	interceptorClient, err := interceptors_client.MewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal(err)
	}

	// create 2nd grpc connection with interceptors
	conn2, err := grpc.NewClient(
		*serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptorClient.Unary()),
		grpc.WithStreamInterceptor(interceptorClient.Stream()),
	)
	if err != nil {
		log.Fatal(err)
	}

	// create laptop service client (this is a grpc client)
	laptopClient := client.NewLaptopClient(conn2)

	// testCreateLaptop(laptopClient)
	// testSearchLaptop(laptopClient)
	// testUploadImage(laptopClient)
	testRatelaptop(laptopClient)
}

func testCreateLaptop(laptopClient *client.LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

func testSearchLaptop(laptopClient *client.LaptopClient) {
	// Create 10 random laptops
	for i := 0; i < 10; i++ {
		laptopClient.CreateLaptop(sample.NewLaptop())
	}

	// Create filter
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	laptopClient.SearchLaptop(filter)
}

func testUploadImage(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	// upload image
	laptopClient.UploadImage(laptop.GetId(), "tmp/stripe.png")
}

func testRatelaptop(laptopClient *client.LaptopClient) {

	// create 3 laptop for rate
	n := 3
	laptopIDs := make([]string, n)
	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	scores := make([]float64, n)
	for {
		fmt.Print("rate laptop [y/n]: ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}
		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}
