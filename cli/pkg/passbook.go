package pkg

import (
	"context"
	"fmt"
	pb "github.com/22Fariz22/passbook/server/proto"
	"log"
	"os"
)

type Acc struct {
	Title string
	Data  string
}

func Register(c pb.UserServiceClient, input *pb.RegisterRequest) error {
	_, err := c.Register(context.Background(), input)
	if err != nil {
		log.Fatal("err", err)
		return err
	}

	return nil
}

func Login(c pb.UserServiceClient, input *pb.LoginRequest) error {
	resp, err := c.Login(context.Background(), input)
	if err != nil {
		log.Fatal("err", err)
		return err
	}

	//пишем в файл session_id
	f, err := os.Create("session.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	l, err := f.WriteString(resp.SessionId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(l, "bytes written successfully")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
