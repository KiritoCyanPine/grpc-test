package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
	"github.com/kiritocyanpine/grpctest/service"
	"google.golang.org/grpc"
)

func main() {

	port := flag.Int("port", 0, "the grpc server port")
	flag.Parse()
	log.Printf("starting server on port %d ", *port)

	laptopserver := service.NewLaptopServer(service.CreateInMemoryLaptopStore())

	grpcserver := grpc.NewServer()
	grpctest.RegisterLaptopServiceServer(grpcserver, laptopserver)

	address := fmt.Sprintf("0.0.0.0:%d", 8080)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listerner port down")
	}

	err = grpcserver.Serve(listener)
	if err != nil {
		panic("could not start server")
	}
}
