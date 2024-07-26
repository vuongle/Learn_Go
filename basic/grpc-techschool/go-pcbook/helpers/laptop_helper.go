package helpers

import (
	"github.com/jinzhu/copier"
	"github.com/vuongle/grpc/pb"
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
