package services

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/vuongle/grpc/helpers"
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
	LaptopStore storages.LaptopStore
	ImageStore  storages.ImageStore
	RatingStore storages.RatingStore
}

func NewLaptopServer(
	laptopStore storages.LaptopStore,
	imageStore storages.ImageStore,
	ratingStore storages.RatingStore,
) *LaptopServer {
	return &LaptopServer{
		LaptopStore: laptopStore,
		ImageStore:  imageStore,
		RatingStore: ratingStore,
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
	err := s.LaptopStore.Save(laptop)
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

	err := s.LaptopStore.Search(
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

func (s *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {

	// Get request from stream
	req, err := stream.Recv()
	if err != nil {
		log.Println("cannot receive image info", err)
		return helpers.LogErr(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	// Get image info from request
	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an image upload for laptop %s with image type %s", laptopID, imageType)

	// find laptop first
	laptop, err := s.LaptopStore.Find(laptopID)
	if err != nil {
		return helpers.LogErr(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return helpers.LogErr(status.Errorf(codes.InvalidArgument, "laptop %s doesn't exist", laptopID))
	}

	// create new byte buffer to save the image
	imageData := bytes.Buffer{}
	imageSize := 0

	// because receive muliplt request from client-strean
	// so must use for loop to receice image data
	for {
		// check context errors(timeout, cancel, ...) before receiving the stream data
		if err := helpers.ContextErr(stream.Context()); err != nil {
			return err
		}

		log.Println("waiting to receive more data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("no more data")
			break
		}
		if err != nil {
			return helpers.LogErr(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		imageSize += size
		if imageSize > helpers.MaxImageSize {
			return helpers.LogErr(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, helpers.MaxImageSize))
		}

		// simulate writing slowly
		time.Sleep(time.Second)

		// write chunk data to buffer
		_, err = imageData.Write(chunk)
		if err != nil {
			return helpers.LogErr(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	// save buffer to disk
	imageID, err := s.ImageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return helpers.LogErr(status.Errorf(codes.Internal, "cannot save image to disk: %v", err))
	}

	// create response and send stream to client
	res := &pb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return helpers.LogErr(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved image with id: %s, size: %d", imageID, imageSize)
	return nil
}

func (s *LaptopServer) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
	// use for loop to receive chunk data from client stream
	for {
		// check context errors(timeout, cancel, ...) before receiving the stream data
		if err := helpers.ContextErr(stream.Context()); err != nil {
			return err
		}

		// receive request from the stream
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("no more data")
			break
		}
		if err != nil {
			return helpers.LogErr(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		laptopID := req.GetLaptopId()
		score := req.GetScore()
		log.Printf("received a rate-laptop request: id = %s, score = %.2f", laptopID, score)

		// find the laptop before rating
		found, err := s.LaptopStore.Find(laptopID)
		if err != nil {
			return helpers.LogErr(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
		}
		if found == nil {
			return helpers.LogErr(status.Errorf(codes.NotFound, "laptop ID %s is not found", laptopID))
		}

		updatedRating, err := s.RatingStore.Add(laptopID, score)
		if err != nil {
			return helpers.LogErr(status.Errorf(codes.Internal, "cannot add rating to the store: %v", err))
		}

		// create response for each chunk data in the stream
		res := &pb.RateLaptopResponse{
			LaptopId:     laptopID,
			RatedCount:   updatedRating.Count,
			AverageScore: updatedRating.Sum / float64(updatedRating.Count),
		}

		// send the response to the client
		err = stream.Send(res)
		if err != nil {
			return helpers.LogErr(status.Errorf(codes.Unknown, "cannot send response: %v", err))
		}
	}

	return nil
}
