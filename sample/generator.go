package sample

import (
	uid "github.com/google/uuid"
	pb "github.com/kiritocyanpine/grpctest/pb/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCPU() *pb.CPU {
	brnd := randomProcessorBrand()

	name := randomCPUName(brnd)

	cores := 2 << randomInt(0, 3)

	threads := cores * 2

	return &pb.CPU{
		Brand:         brnd,
		Name:          name,
		NumberCores:   uint32(cores),
		NumberThreads: uint32(threads),
		MinGhz:        randomFloat64(1.5, 3.9),
		MaxGhz:        randomFloat64(3.2, 4.9),
	}
}

func NewGPU() *pb.GPU {
	brnd := randomGPUBrand()

	name := randomGPUName(brnd)
	return &pb.GPU{
		Brand:         brnd,
		Name:          name,
		NumberThreads: 4500,
		MinGhz:        randomFloat64(1.5, 3.9),
		MaxGhz:        randomFloat64(3.2, 4.9),
		Memory: &pb.Memory{
			Value: uint64(256 * randomInt(1, 8)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}

func NewStorage() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: 1,
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewMemory() *pb.Memory {
	return &pb.Memory{
		Value: 2 << randomInt(0, 3),
		Unit:  pb.Memory_GIGABYTE,
	}
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	return &pb.Laptop{
		Id:       uid.NewString(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Memory:   NewMemory(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewStorage()},
		Weight: &pb.Laptop_WeightKg{
			WeightKg: 2,
		},
		Price: randomFloat64(1250, 2690),

		UpdatedAt: timestamppb.Now(),
	}
}
