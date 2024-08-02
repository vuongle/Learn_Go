package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/vuongle/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopClient struct {
	service pb.LaptopServiceClient
}

func NewLaptopClient(cc *grpc.ClientConn) *LaptopClient {
	return &LaptopClient{service: pb.NewLaptopServiceClient(cc)}
}

func (laptopClient *LaptopClient) CreateLaptop(laptop *pb.Laptop) {

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.service.CreateLaptop(ctx, req)
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

func (laptopClient *LaptopClient) SearchLaptop(filter *pb.Filter) {
	log.Printf("search filter: %v", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}

	stream, err := laptopClient.service.SearchLaptop(ctx, req)
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

func (laptopClient *LaptopClient) UploadImage(laptopID string, imagePath string) {
	file, err := os.Open(imagePath) // open image file on client side
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create client stream to upload, not send yet
	stream, err := laptopClient.service.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: filepath.Ext(imagePath), // get image type
			},
		},
	}

	// send image info (only ID and type) first, not send data yet
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024) // 1KB buffer
	for {
		// read 1KB for each time
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		// send image data (chunk data)
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk data: ", err, stream.RecvMsg(nil))
		}
	}

	// close stream and wait for response.
	// Once the clients close the stream -> server receive end of stream (io.EOF in laptop_server.go)
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded successfully: %s", res.GetId())
}

func (laptopClient *LaptopClient) RateLaptop(
	laptopIDs []string,
	scores []float64,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create a stream for rating, not send yet
	stream, err := laptopClient.service.RateLaptop(ctx)
	if err != nil {
		return fmt.Errorf("cannot create stream: %v", err)
	}

	// create a channel so that the client can can send and receive concurrently
	// waitResponse:
	//    receive an err if there is any err
	//    or
	//    nil if the client receives all messages successfully
	waitResponse := make(chan error)
	// start a goroutine to receive response while the client still send to the server
	go func() {
		for {
			// receive response from the server
			res, err := stream.Recv()
			// if EOF -> no more data in the stream -> send "nil" to the channel
			if err == io.EOF {
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive response: %v", err)
				return
			}

			log.Println("received response: ", res)
		}
	}()

	// Send message to server
	for i, laptopID := range laptopIDs {
		req := &pb.RateLaptopRequest{
			LaptopId: laptopID,
			Score:    scores[i],
		}

		err := stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send stream: %v - %v", err, stream.RecvMsg(nil))
		}

		log.Println("sent request: ", req)
	}

	// close the stream
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close stream: %v", err)
	}

	// read the channel
	err = <-waitResponse
	return err
}
