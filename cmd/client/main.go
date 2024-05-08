package main

import (
	"context"
	"flag"
	"log"

	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
	"github.com/kiritocyanpine/grpctest/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	serveradd := flag.String("address", "", "the address of the server")
	flag.Parse()
	log.Println(" dialing to server")

	conn, err := grpc.Dial(*serveradd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to server")
	}

	lclient := grpctest.NewLaptopServiceClient(conn)

	lap := sample.NewLaptop()

	req := &grpctest.CreateLaptopRequest{
		Laptop: lap,
	}

	res, err := lclient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("already exists")
		} else {
			log.Fatal("cannot create")
		}

		return
	}

	req2 := &grpctest.GetLaptopListRequest{}

	res2, err := lclient.GetLaptop(context.Background(), req2)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("already exists")
		} else {
			log.Fatal("cannot create")
		}

		return
	}

	log.Printf("Laptop created with ID :%s", res.Id)

	log.Printf("Laptops here : %v", res2.Laptops)

}
