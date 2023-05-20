package service

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/22Fariz22/passbook/server/config"
	mockSessUC "github.com/22Fariz22/passbook/server/internal/session/mock"
	"github.com/22Fariz22/passbook/server/internal/user/mock"
	"github.com/22Fariz22/passbook/server/pkg/logger"
	userService "github.com/22Fariz22/passbook/server/proto"
)

func TestUsersService_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userUC := mock.NewMockUserUseCase(ctrl)
	sessUC := mockSessUC.NewMockSessionUseCase(ctrl)
	apiLogger := logger.NewAPILogger(nil)
	authServerGRPC := NewAuthServerGRPC(apiLogger, nil, userUC, sessUC)

	reqValue := &userService.RegisterRequest{
		Login:    "email@gmail.com",
		Password: "Password",
	}

	t.Run("Register", func(t *testing.T) {
		t.Parallel()
		userID := uuid.New()
		user := &entity.User{
			UserID:   userID,
			Login:    reqValue.Login,
			Password: reqValue.Password,
		}

		userUC.EXPECT().Register(gomock.Any(), gomock.Any()).Return(user, nil)

		response, err := authServerGRPC.Register(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, reqValue.Login, response.User.Login)
	})
}

func TestUsersService_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userUC := mock.NewMockUserUseCase(ctrl)
	sessUC := mockSessUC.NewMockSessionUseCase(ctrl)
	apiLogger := logger.NewAPILogger(nil)
	cfg := &config.Config{Session: config.Session{
		Expire: 10,
	}}
	authServerGRPC := NewAuthServerGRPC(apiLogger, cfg, userUC, sessUC)

	reqValue := &userService.LoginRequest{
		Login:    "email@gmail.com",
		Password: "Password",
	}

	t.Run("Login", func(t *testing.T) {
		t.Parallel()
		userID := uuid.New()
		session := "session"
		user := &entity.User{
			UserID:   userID,
			Login:    "email@gmail.com",
			Password: "Password",
		}

		userUC.EXPECT().Login(gomock.Any(), reqValue.Login, reqValue.Password).Return(user, nil)
		sessUC.EXPECT().CreateSession(gomock.Any(), &entity.Session{
			UserID: user.UserID,
		}, cfg.Session.Expire).Return(session, nil)

		response, err := authServerGRPC.Login(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, reqValue.Login, response.User.Login)
	})
}

func TestUsersService_FindByLogin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userUC := mock.NewMockUserUseCase(ctrl)
	sessUC := mockSessUC.NewMockSessionUseCase(ctrl)
	apiLogger := logger.NewAPILogger(nil)
	cfg := &config.Config{Session: config.Session{
		Expire: 10,
	}}
	authServerGRPC := NewAuthServerGRPC(apiLogger, cfg, userUC, sessUC)

	reqValue := &userService.FindByLoginRequest{
		Login: "email@gmail.com",
	}

	t.Run("FindByLogin", func(t *testing.T) {
		t.Parallel()
		userID := uuid.New()
		user := &entity.User{
			UserID:   userID,
			Login:    "email@gmail.com",
			Password: "Password",
		}

		userUC.EXPECT().FindByLogin(gomock.Any(), reqValue.Login).Return(user, nil)

		response, err := authServerGRPC.FindByLogin(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, reqValue.Login, response.User.Login)
	})
}
