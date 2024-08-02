package services_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/sample"
	"github.com/vuongle/grpc/serializer"
	"github.com/vuongle/grpc/services"
	"github.com/vuongle/grpc/storages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := storages.NewInMemoryLaptopStore()
	imageStore := storages.NewDiskImageStore("img")
	ratingStore := storages.NewInMemoryRatingStore()

	laptopServer, serverAddr := startGRPCServer(t, laptopStore, imageStore, ratingStore)
	laptopClient := newTestLaptopClient(t, serverAddr)

	laptop := sample.NewLaptop()
	laptopID := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, laptopID, res.Id)

	// check that the laptop is saved to the store correctly
	saved, err := laptopServer.LaptopStore.Find(laptopID)
	require.NoError(t, err)
	require.NotNil(t, saved)
	requireSameLaptop(t, laptop, saved)
}

func startGRPCServer(
	t *testing.T,
	laptopStore storages.LaptopStore,
	imageStore storages.ImageStore,
	ratingStore storages.RatingStore,
) (*services.LaptopServer, string) {
	laptopServer := services.NewLaptopServer(laptopStore, imageStore, ratingStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", ":0") //":0" will use a random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener) // Serve() is a blocking call so we need to run it in a separate goroutine

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddr string) pb.LaptopServiceClient {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	// require.Equal(t, laptop1, laptop2)
	// can not use above code becausepb.Laptop have generated fields (state, sizeCache)
	// that can not be compated
	// Therefore, convert to json

	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
