package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	// get server address from command line
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dialing server: %s", *serverAddress)

	// create grpc connection
	conn, err := grpc.NewClient(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// create laptop service client (this is a grpc client)
	client := pb.NewLaptopServiceClient(conn)

	// Create 10 random laptops
	for i := 0; i < 10; i++ {
		createLaptop(client)
	}

	// Create filter
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	searchLaptop(client, filter)
}

func createLaptop(client pb.LaptopServiceClient) {

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}

		return
	}

	log.Printf("created laptop with id: %s", res.Id)
}

func searchLaptop(client pb.LaptopServiceClient, filter *pb.Filter) {
	log.Printf("search filter: %v", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}

	stream, err := client.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}

	for {
		res, err := stream.Recv()

		// end of stream -> receive end of stream
		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		foundLaptop := res.GetLaptop()
		log.Printf("- found: %s", foundLaptop.GetId())
		log.Printf("    + brand: %s", foundLaptop.GetBrand())
		log.Printf("    + name: %s", foundLaptop.GetName())
		log.Printf("    + cpu cores: %d", foundLaptop.GetCpu().GetNumberCores())
		log.Printf("    + cpu min ghz: %f", foundLaptop.GetCpu().GetMinGhz())
		log.Printf("    + ram: %d", foundLaptop.GetMemory().GetValue())
		log.Printf("    + price: %f", foundLaptop.GetPriceUsd())
	}
}
