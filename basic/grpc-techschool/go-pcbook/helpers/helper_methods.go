package helpers

import (
	"context"
	"log"

	"github.com/jinzhu/copier"
	"github.com/vuongle/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if ToBit(laptop.GetMemory()) < ToBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func ToBit(m *pb.Memory) uint64 {
	value := m.GetValue()

	switch m.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value * 8
	case pb.Memory_KILOBYTE:
		return value * 1024 * 8
	case pb.Memory_MEGABYTE:
		return value * 1024 * 1024 * 8
	case pb.Memory_GIGABYTE:
		return value * 1024 * 1024 * 1024 * 8
	case pb.Memory_TERABYTE:
		return value * 1024 * 1024 * 1024 * 1024 * 8
	default:
		return 0
	}
}

func DeepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, err
	}

	return other, nil
}

func LogErr(err error) error {
	if err != nil {
		log.Println(err)
	}

	return err
}

func ContextErr(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return LogErr(status.Errorf(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return LogErr(status.Errorf(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func AccessibleRoles() map[string][]string {
	const laptopServicePath = "/grpcservice.laptop.LaptopService/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}
