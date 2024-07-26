package services

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/storages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// / This struct is gRPC server API that implements the generated LaptopServiceServer interface(in laptop_service_grpc.pb.go)
// / All clients (Go, web, mobile that support gRPC) will call this server API.
// / The LaptopServiceServer interface is generated based on services defined in laptop_service.proto
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store storages.LaptopStore
}

func NewLaptopServer(store storages.LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}

// / Impletement CreateLaptop() in the LaptopServiceServer interface (in laptop_service_grpc.pb.go)
func (s *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {

	// Get data from request
	laptop := req.GetLaptop()
	log.Printf("receive laptop with id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		// check if it's valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// simulate slow network or heavy processing
	//time.Sleep(6 * time.Second)

	// check timeout or cancel error to prevent the server still save the data while client has caneled or timed out
	// if not check and handle error here, the code still runs and save the data even if client has caneled or timed out
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("request is timeout")
		return nil, status.Errorf(codes.DeadlineExceeded, "request is timeout")
	}
	if ctx.Err() == context.Canceled {
		log.Println("request is canceled")
		return nil, status.Errorf(codes.Canceled, "request is canceled")
	}

	// Save data
	err := s.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, storages.ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}

// SearchLaptop is a server-streaming RPC to search for laptops that match the provided filters.
//
// Parameters:
// - req: The search request containing the search criteria.
// - stream: The stream used to send the search results to the client.
//
// Returns:
// - error: An error if there was a problem performing the search.
func (s *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer) error {

	// Get filter from request
	filter := req.GetFilter()
	log.Printf("receive a search laptop request with filter: %v", filter)

	err := s.Store.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{
				Laptop: laptop,
			}

			// send stream to client
			err := stream.Send(res)
			if err != nil {
				return err
			}

			log.Printf("sent laptop with id: %s", laptop.GetId())
			return nil
		})

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
}
