package service

import (
	"context"
	"errors"
	"log"

	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/22Fariz22/passbook/server/pkg/grpcerrors"
	"github.com/22Fariz22/passbook/server/pkg/utils"
	userService "github.com/22Fariz22/passbook/server/proto"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Register new user
func (u *usersService) Register(ctx context.Context, r *userService.RegisterRequest) (*userService.RegisterResponse, error) {
	user, err := u.registerReqToUserModel(r)
	if err != nil {
		u.logger.Errorf("registerReqToUserModel: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "registerReqToUserModel: %v", err)
	}

	if err := utils.ValidateStruct(ctx, user); err != nil {
		u.logger.Errorf("ValidateStruct: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "ValidateStruct: %v", err)
	}

	createdUser, err := u.userUC.Register(ctx, user)
	if err != nil {
		u.logger.Errorf("userUC.Register: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "Register: %v", err)
	}

	return &userService.RegisterResponse{User: u.userModelToProto(createdUser)}, nil
}

// Login user with email and password
func (u *usersService) Login(ctx context.Context, r *userService.LoginRequest) (*userService.LoginResponse, error) {
	login := r.GetLogin()

	user, err := u.userUC.Login(ctx, login, r.GetPassword())
	if err != nil {
		u.logger.Errorf("userUC.Login: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "Login: %v", err)
	}

	session, err := u.sessUC.CreateSession(ctx, &entity.Session{
		UserID: user.UserID,
	}, u.cfg.Session.Expire)
	if err != nil {
		u.logger.Errorf("sessUC.CreateSession: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "sessUC.CreateSession: %v", err)
	}

	return &userService.LoginResponse{User: u.userModelToProto(user), SessionId: session}, err
}

// FindByLogin Find user by login
func (u *usersService) FindByLogin(ctx context.Context, r *userService.FindByLoginRequest) (*userService.FindByLoginResponse, error) {
	login := r.GetLogin()

	user, err := u.userUC.FindByLogin(ctx, login)
	if err != nil {
		u.logger.Errorf("userUC.FindByLogin: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "userUC.FindByLogin: %v", err)
	}

	return &userService.FindByLoginResponse{User: u.userModelToProto(user)}, err
}

// FindByID Find user by uuid
func (u *usersService) FindByID(ctx context.Context, r *userService.FindByIDRequest) (*userService.FindByIDResponse, error) {
	userUUID, err := uuid.Parse(r.GetUuid())
	if err != nil {
		u.logger.Errorf("uuid.Parse: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "uuid.Parse: %v", err)
	}

	user, err := u.userUC.FindByID(ctx, userUUID)
	if err != nil {
		u.logger.Errorf("userUC.FindByID: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "userUC.FindByID: %v", err)
	}

	return &userService.FindByIDResponse{User: u.userModelToProto(user)}, nil
}

// GetMe Get session id from, ctx metadata, find user by uuid and returns it
func (u *usersService) GetMe(ctx context.Context, r *userService.GetMeRequest) (*userService.GetMeResponse, error) {
	sessID, err := u.getSessionIDFromCtx(ctx)
	if err != nil {
		u.logger.Errorf("getSessionIDFromCtx: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "sessUC.getSessionIDFromCtx: %v", err)
	}

	session, err := u.sessUC.GetSessionByID(ctx, sessID)
	if err != nil {
		u.logger.Errorf("sessUC.GetSessionByID: %v", err)
		if errors.Is(err, redis.Nil) {
			return nil, status.Errorf(codes.NotFound, "sessUC.GetSessionByID: %v", grpcerrors.ErrNotFound)
		}
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "sessUC.GetSessionByID: %v", err)
	}

	user, err := u.userUC.FindByID(ctx, session.UserID)
	if err != nil {
		u.logger.Errorf("userUC.FindByID: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "userUC.FindByID: %v", err)
	}

	return &userService.GetMeResponse{User: u.userModelToProto(user)}, nil
}

// Logout user, delete current session
func (u *usersService) Logout(ctx context.Context, request *userService.LogoutRequest) (*userService.LogoutResponse, error) {
	sessID, err := u.getSessionIDFromCtx(ctx)
	if err != nil {
		u.logger.Errorf("getSessionIDFromCtx: %v", err)
		return nil, err
	}
	if err := u.sessUC.DeleteByID(ctx, sessID); err != nil {
		log.Println("here err u.sessUC.DeleteByID.")
		u.logger.Errorf("sessUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "sessUC.DeleteByID: %v", err)
	}

	return &userService.LogoutResponse{}, nil
}

// AddAccount save account data
func (u *usersService) AddAccount(ctx context.Context, request *userService.AddAccountRequest) (*userService.AddAccountResponse, error) {
	session, err := checkSessionAndGetUserID(ctx, u)
	if err != nil {
		return nil, err
	}

	err = u.userUC.AddAccount(ctx, session.UserID.String(), request)
	if err != nil {
		return &userService.AddAccountResponse{}, err
	}
	if request.Login == "" || request.Password == "" {
		return &userService.AddAccountResponse{}, errors.New("failed to add, because your account is empty")
	}

	return &userService.AddAccountResponse{}, nil
}

// AddText save text data
func (u *usersService) AddText(ctx context.Context, request *userService.AddTextRequest) (*userService.AddTextResponse, error) {
	session, err := checkSessionAndGetUserID(ctx, u)
	if err != nil {
		return nil, err
	}

	err = u.userUC.AddText(ctx, session.UserID.String(), request)
	if err != nil {
		return &userService.AddTextResponse{}, err
	}
	if request.Data == "" {
		return &userService.AddTextResponse{}, errors.New("failed to add text, because your data is empty")
	}

	return &userService.AddTextResponse{}, nil
}

// AddBinary save binary data
func (u *usersService) AddBinary(ctx context.Context, request *userService.AddBinaryRequest) (*userService.AddBinaryResponse, error) {
	session, err := checkSessionAndGetUserID(ctx, u)
	if err != nil {
		return nil, err
	}

	err = u.userUC.AddBinary(ctx, session.UserID.String(), request)
	if err != nil {
		return &userService.AddBinaryResponse{}, err
	}
	if len(request.Data) == 0 {
		return &userService.AddBinaryResponse{}, errors.New("failed to add binary, because your data is empty")
	}

	return &userService.AddBinaryResponse{}, nil
}

// AddCard save to card data
func (u *usersService) AddCard(ctx context.Context, request *userService.AddCardRequest) (*userService.AddCardResponse, error) {
	session, err := checkSessionAndGetUserID(ctx, u)
	if err != nil {
		return nil, err
	}

	err = u.userUC.AddCard(ctx, session.UserID.String(), request)
	if err != nil {
		return &userService.AddCardResponse{}, err
	}
	if request.CardNumber == "" || request.CvcCode == "" || request.DateExp == "" {
		return &userService.AddCardResponse{}, errors.New("failed to add card, because your number ,cvc code or data exp is empty")
	}

	return &userService.AddCardResponse{}, nil
}

// GetByTitle get user's data by title
func (u *usersService) GetByTitle(ctx context.Context, request *userService.GetByTitleRequest) (*userService.GetByTitleResponse, error) {
	session, err := checkSessionAndGetUserID(ctx, u)
	if err != nil {
		return nil, err
	}

	data, err := u.userUC.GetByTitle(ctx, session.UserID.String(), request)
	if err != nil {
		return &userService.GetByTitleResponse{}, err
	}

	return &userService.GetByTitleResponse{Data: data}, err
}

// GetFullList get all user's data
func (u *usersService) GetFullList(ctx context.Context, request *userService.GetFullListRequest) (*userService.GetFullListResponse, error) {
	session, err := checkSessionAndGetUserID(ctx, u)
	if err != nil {
		return nil, err
	}

	data, err := u.userUC.GetFullList(ctx, session.UserID.String())
	if err != nil {
		return &userService.GetFullListResponse{}, err
	}

	return &userService.GetFullListResponse{Data: data}, nil
}

func (u *usersService) registerReqToUserModel(r *userService.RegisterRequest) (*entity.User, error) {
	candidate := &entity.User{
		Login:    r.GetLogin(),
		Password: r.GetPassword(),
	}

	if err := candidate.PrepareCreate(); err != nil {
		return nil, err
	}

	return candidate, nil
}

func (u *usersService) userModelToProto(user *entity.User) *userService.User {
	userProto := &userService.User{
		Uuid:     user.UserID.String(),
		Login:    user.Login,
		Password: user.Password,
	}
	return userProto
}

func (u *usersService) getSessionIDFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext: %v", grpcerrors.ErrNoCtxMetaData)
	}

	sessionID := md.Get("session_id")
	if sessionID[0] == "" {
		return "", status.Errorf(codes.PermissionDenied, "md.Get sessionId: %v", grpcerrors.ErrInvalidSessionID)
	}

	return sessionID[0], nil
}

func checkSessionAndGetUserID(ctx context.Context, u *usersService) (*entity.Session, error) {
	sessID, err := u.getSessionIDFromCtx(ctx)
	if err != nil {
		u.logger.Errorf("getSessionIDFromCtx: %v", err)
		return nil, err
	}

	session, err := u.sessUC.GetSessionByID(ctx, sessID)
	if err != nil {
		u.logger.Errorf("sessUC.GetSessionByID: %v", err)
		if errors.Is(err, redis.Nil) {
			return nil, status.Errorf(codes.NotFound, "sessUC.GetSessionByID: %v", grpcerrors.ErrNotFound)
		}
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "sessUC.GetSessionByID: %v", err)
	}
	return session, nil
}
