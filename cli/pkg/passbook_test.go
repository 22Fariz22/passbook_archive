package pkg

import (
	"context"
	"errors"
	pb "github.com/22Fariz22/passbook/server/proto"
	"testing"
)

//type IConn interface {
//	ConnGRPCServer() pb.UserServiceClient
//}

func iConnGRPCServer() *pb.UserServiceClient {
	//server := grpc.NewServer()

	c := pb.UserServiceClient(nil)
	return &c
}

func TestRegister(t *testing.T) {
	ctx := context.Background()

	client := ConnGRPCServer() //замокать!

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
				out: pb.RegisterResponse{User: &pb.User{
					Uuid:     "",
					Login:    "",
					Password: "",
				}},
				err: errors.New(""),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.Register(ctx, &tt.in)
			if err != nil {
				//t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}

}
