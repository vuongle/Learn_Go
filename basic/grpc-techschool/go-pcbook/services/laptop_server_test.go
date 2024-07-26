package services_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/sample"
	"github.com/vuongle/grpc/services"
	"github.com/vuongle/grpc/storages"
	"google.golang.org/grpc/codes"
)

func TestServiceCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-id"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := storages.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)

	// Define test cases
	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  storages.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  storages.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  storages.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "fail_invalid_id",
			laptop: laptopInvalidID,
			store:  storages.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "fail_duplicate_id",
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			laptopService := services.NewLaptopServer(tc.store)

			res, err := laptopService.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, res)
			}
		})
	}
}
