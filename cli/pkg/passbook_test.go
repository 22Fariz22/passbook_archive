package pkg

import (
	"context"
	"testing"

	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/22Fariz22/passbook/server/proto/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	t.Parallel()

	//client := ConnGRPCServer() // реальный сервер

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockUserServiceServer(ctrl) //NewMockUserServiceClient(ctrl)

	in := pb.RegisterRequest{
		Login:    "Leo",
		Password: "qwerty",
	}

	client.EXPECT().Register(ctx, &in).Return(&pb.RegisterResponse{}, nil)

	_, err := client.Register(ctx, &in)
	require.NoError(t, err)
}
