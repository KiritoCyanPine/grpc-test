package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	grpctest.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func NewLaptopServer(s LaptopStore) *LaptopServer {
	return &LaptopServer{grpctest.UnimplementedLaptopServiceServer{}, s}
}

// CreateLaptop is a unary RPC to create a Laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *grpctest.CreateLaptopRequest,
) (*grpctest.CreateLaptopResponse, error) {

	laptop := req.GetLaptop()

	if len(laptop.Id) > 0 {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid laptop uuid: %s", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, fmt.Errorf("unable to generate uuid: %s", err)
		}
		laptop.Id = id.String()
	}

	// save the Laptop to DB
	if err := server.Store.Save(laptop); err != nil {
		code := codes.Internal

		if errors.Is(err, ErrAlreadyExist) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop: %v", err)
	}

	res := &grpctest.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}

func (server *LaptopServer) GetLaptop(
	ctx context.Context,
	req *grpctest.GetLaptopListRequest,
) (*grpctest.GetLaptopListResponse, error) {
	cache := server.Store.GetAll()

	res := &grpctest.GetLaptopListResponse{
		Laptops: cache,
	}

	return res, nil
}
