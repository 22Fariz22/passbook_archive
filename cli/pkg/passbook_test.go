package pkg

import (
	"context"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/22Fariz22/passbook/server/proto/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	t.Parallel()

	//client := ConnGRPCServer() // реальный сервер

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockUserServiceServer(ctrl) //NewMockUserServiceClient(ctrl)

	type expectation struct {
		out pb.RegisterResponse
		err error
	}

	var req *pb.RegisterRequest = new(pb.RegisterRequest)
	req.Login = "Leo9"
	req.Password = "qwer"

	tests := map[string]struct {
		in       pb.RegisterRequest
		expected expectation
	}{
		"Must_Success": {
			in: *req,
			expected: expectation{
				out: pb.RegisterResponse{},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			client.EXPECT().Register(ctx, &tt.in).Return(&pb.RegisterResponse{}, nil)

			_, err := client.Register(ctx, &tt.in)
			require.NoError(t, err)
			if err != nil {
				t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
			}
		})
	}

}
