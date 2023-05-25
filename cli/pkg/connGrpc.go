package pkg

import (
	pb "github.com/22Fariz22/passbook/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// ConnGRPCServer соединят с сервером
func ConnGRPCServer() pb.UserServiceClient {
	conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("conn:", err)
		return nil
	}

	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения
	c := pb.NewUserServiceClient(conn)

	return c
}
