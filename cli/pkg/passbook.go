package pkg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	pb "github.com/22Fariz22/passbook/server/proto"
	"google.golang.org/grpc/metadata"
)

// Register new user
func Register(c pb.UserServiceClient, input *pb.RegisterRequest) error {
	_, err := c.Register(context.Background(), input)
	if err != nil {
		fmt.Println("failed to register")
		return err
	}
	fmt.Println("register is successful")
	return nil
}

// Login user with email and password
func Login(c pb.UserServiceClient, input *pb.LoginRequest) error {
	resp, err := c.Login(context.Background(), input)
	if err != nil {
		fmt.Println("failed to login")
		return err
	}

	//пишем в файл session_id
	f, err := os.Create("session.txt")
	if err != nil {
		fmt.Println("failed to login")
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
		fmt.Println("failed to logout")
		return err
	}

	_, err = c.Logout(ctx, &pb.LogoutRequest{})
	if err != nil {
		fmt.Println("failed to logout")
		return err
	}
	fmt.Println("logout is successful")
	return nil
}

// GetMe посмотреть свою сессию
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

// AddAccount  сохранить данные аккаунта
func AddAccount(c pb.UserServiceClient, input *pb.AddAccountRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return err
	}

	_, err = c.AddAccount(ctx, input)
	if err != nil {
		fmt.Println("failed to add account")
		return err
	}
	fmt.Println("account added")
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
		fmt.Println("failed to add text")
		return err
	}
	fmt.Println("text added")
	return nil
}

// AddBinary add text data
func AddBinary(c pb.UserServiceClient, input *pb.AddBinaryRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		return err
	}

	_, err = c.AddBinary(ctx, input)
	if err != nil {
		fmt.Println("failed to add binary")
		return err
	}
	fmt.Println("binary added")
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
		fmt.Println("failed to add card")
		return err
	}

	fmt.Println("card added")
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
	data, err := os.ReadFile("session.txt")
	if err != nil {
		log.Println("err in ioutil.ReadFile:", err)
		return nil, err
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return ctx, nil
}
