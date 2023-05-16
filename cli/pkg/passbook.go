package pkg

import (
	"context"
	"fmt"
	pb "github.com/22Fariz22/passbook/server/proto"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
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
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}

	return nil
}

func Logout(c pb.UserServiceClient, input *pb.LogoutRequest) error {
	data, err := ioutil.ReadFile("session.txt")
	if err != nil {
		log.Fatal("err in ioutil.ReadFile", err)
		return nil
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err = c.Logout(ctx, &pb.LogoutRequest{})
	if err != nil {
		log.Fatal("err in c.Logout", err)
		return err
	}

	return nil
}

func AddAccount(c pb.UserServiceClient, input *pb.AddAccountRequest) error {
	data, err := ioutil.ReadFile("session.txt")
	if err != nil {
		log.Println("err in ioutil.ReadFile:", err)
		return err
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err = c.AddAccount(ctx, input)
	if err != nil {
		//add err
	}

	return nil
}

func AddText(c pb.UserServiceClient, input *pb.AddTextRequest) error {
	data, err := ioutil.ReadFile("session.txt")
	if err != nil {
		log.Println("err in ioutil.ReadFile:", err)
		return err
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err = c.AddText(ctx, input)
	if err != nil {
		//add err
	}

	return nil
}

func AddCard(c pb.UserServiceClient, input *pb.AddCardRequest) error {
	data, err := ioutil.ReadFile("session.txt")
	if err != nil {
		log.Println("err in ioutil.ReadFile:", err)
		return err
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err = c.AddCard(ctx, input)
	if err != nil {
		//add err
	}

	return nil
}
