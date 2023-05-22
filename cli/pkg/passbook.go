package pkg

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/22Fariz22/passbook/server/proto"
	"google.golang.org/grpc/metadata"
)

// Register new user
func Register(c pb.UserServiceClient, input *pb.RegisterRequest) error {
	_, err := c.Register(context.Background(), input)
	if err != nil {
		log.Fatal("err", err)
		return err
	}

	return nil
}

// Login user with email and password
func Login(c pb.UserServiceClient, input *pb.LoginRequest) error {
	resp, err := c.Login(context.Background(), input)
	if err != nil {
		log.Fatal("err", err)
		return err
	}

	//пишем в файл session_id
	f, err := os.Create("session.txt")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = f.WriteString(resp.SessionId)
	if err != nil {
		err = errors.New("возможно вы еще не зарегистрировались")
		log.Println(err)
		return err
	}

	fmt.Println("login successfully")

	return nil
}

// Logout выход из системы
func Logout(c pb.UserServiceClient, input *pb.LogoutRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return err
	}

	_, err = c.Logout(ctx, &pb.LogoutRequest{})
	if err != nil {
		log.Fatal("err in c.Logout", err)
		return err
	}

	return nil
}

func GetMe(c pb.UserServiceClient, input *pb.GetMeRequest) (*pb.GetMeResponse, error) {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return &pb.GetMeResponse{}, err
	}

	user, err := c.GetMe(ctx, input)
	if err != nil {
		return &pb.GetMeResponse{}, err
	}

	return user, nil
}

// AddAccount вызов rpc AddAccount -> добавить в хранилище сведдения об аккаунте
func AddAccount(c pb.UserServiceClient, input *pb.AddAccountRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return err
	}

	_, err = c.AddAccount(ctx, input)
	if err != nil {
		//add err
	}

	return nil
}

// AddText add text data
func AddText(c pb.UserServiceClient, input *pb.AddTextRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return err
	}

	_, err = c.AddText(ctx, input)
	if err != nil {
		//add err
	}

	return nil
}

// AddCard add card data
func AddCard(c pb.UserServiceClient, input *pb.AddCardRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return err
	}

	_, err = c.AddCard(ctx, input)
	if err != nil {
		//add err
	}

	return nil
}

// GetByTitle find data by title
func GetByTitle(c pb.UserServiceClient, input *pb.GetByTitleRequest) (*pb.GetByTitleResponse, error) {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return nil, err
	}

	res, err := c.GetByTitle(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetFullList find all type of data
func GetFullList(c pb.UserServiceClient, input *pb.GetFullListRequest) (*pb.GetFullListResponse, error) {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return nil, err
	}

	res, err := c.GetFullList(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSessionAndPutInMD читает файл session_id, ищет сессию ,вставляет session_id в метаданные и возвращает context
func GetSessionAndPutInMD() (context.Context, error) {
	data, err := ioutil.ReadFile("session.txt")
	if err != nil {
		log.Println("err in ioutil.ReadFile:", err)
		return nil, err
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return ctx, nil
}
