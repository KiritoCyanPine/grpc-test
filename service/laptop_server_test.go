package service_test

import (
	"context"
	"testing"

	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
	"github.com/kiritocyanpine/grpctest/sample"
	"github.com/kiritocyanpine/grpctest/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopnoid := sample.NewLaptop()
	laptopnoid.Id = ""

	laptopinvalidid := sample.NewLaptop()
	laptopinvalidid.Id = "invalid_uuid"

	duplicateLaptop := sample.NewLaptop()

	storeduplicateId := service.CreateInMemoryLaptopStore()
	err := storeduplicateId.Save(duplicateLaptop)
	require.NoError(t, err)

	testcases := []struct {
		name   string
		laptop *grpctest.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success with id",
			laptop: sample.NewLaptop(),
			store:  service.CreateInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success with no id",
			laptop: laptopnoid,
			store:  service.CreateInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure invalid id",
			laptop: laptopinvalidid,
			store:  service.CreateInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure duplicate id",
			laptop: duplicateLaptop,
			store:  storeduplicateId,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testcases {
		tc := testcases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &grpctest.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer(tc.store)

			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)

				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				s, ok := status.FromError(err)
				require.True(t, ok)

				require.Equal(t, s.Code(), tc.code)
			}
		})
	}
}
