package service_test

import (
	"context"
	"encoding/json"
	"net"
	"testing"

	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
	"github.com/kiritocyanpine/grpctest/sample"
	"github.com/kiritocyanpine/grpctest/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientTest(t *testing.T) {
	t.Parallel()

	server, uri := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, uri)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id

	req := &grpctest.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)

	require.Equal(t, res.Id, expectedID)

	val, err := server.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, val)

	compareLaptops(t, val, laptop)

	// laptop dosent exist
	laptopNew := sample.NewLaptop()
	val, err = server.Store.Find(laptopNew.Id)

	require.Error(t, err)
	require.ErrorIs(t, err, service.ErrNotFound)
	require.Nil(t, val)
}

func compareLaptops(t *testing.T, lap1, lap2 *grpctest.Laptop) {
	val1, err := json.Marshal(lap1)
	require.NoError(t, err)

	val2, err := json.Marshal(lap2)
	require.NoError(t, err)

	require.Equal(t, val1, val2)
}

func newTestLaptopClient(t *testing.T, address string) grpctest.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	return grpctest.NewLaptopServiceClient(conn)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopserver := service.NewLaptopServer(service.CreateInMemoryLaptopStore())

	grpcserver := grpc.NewServer()
	grpctest.RegisterLaptopServiceServer(grpcserver, laptopserver)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcserver.Serve(listener)

	return laptopserver, listener.Addr().String()
}
