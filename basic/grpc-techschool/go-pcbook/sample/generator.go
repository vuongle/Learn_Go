package sample

import (
	"github.com/vuongle/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewCPU() *pb.CPU {
	brand := randomCPUbrand()
	name := randomCPUName(brand)
	cores := randomInt(2, 8)
	threads := randomInt(cores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5)

	return &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(cores),
		NumberThreads: uint32(threads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}

// NewRAM returns a new sample RAM
func NewRAM() *pb.Memory {
	memGB := randomInt(4, 64)

	ram := &pb.Memory{
		Value: uint64(memGB),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *pb.Storage {
	memGB := randomInt(128, 1024)

	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(memGB),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *pb.Storage {
	memTB := randomInt(1, 6)

	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(memTB),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Memory:   NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}
