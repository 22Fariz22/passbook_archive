package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/22Fariz22/passbook/server/internal/user/mock"
	"github.com/22Fariz22/passbook/server/pkg/logger"
)

func TestUserUseCase_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userPGRepository := mock.NewMockUserPGRepository(ctrl)
	userRedisRepository := mock.NewMockUserRedisRepository(ctrl)
	apiLogger := logger.NewAPILogger(nil)
	userUC := NewUserUseCase(apiLogger, userPGRepository, userRedisRepository)

	userID := uuid.New()
	mockUser := &entity.User{
		Login:    "email@gmail.com",
		Password: "123456",
	}

	ctx := context.Background()

	userPGRepository.EXPECT().FindByLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	userPGRepository.EXPECT().Create(gomock.Any(), mockUser).Return(&entity.User{
		UserID:   userID,
		Login:    "email@gmail.com",
		Password: "123456",
	}, nil)

	createdUser, err := userUC.Register(ctx, mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.Equal(t, createdUser.UserID, userID)
}

func TestUserUseCase_FindByEmail(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userPGRepository := mock.NewMockUserPGRepository(ctrl)
	userRedisRepository := mock.NewMockUserRedisRepository(ctrl)
	apiLogger := logger.NewAPILogger(nil)
	userUC := NewUserUseCase(apiLogger, userPGRepository, userRedisRepository)

	userID := uuid.New()
	mockUser := &entity.User{
		UserID:   userID,
		Login:    "email@gmail.com",
		Password: "123456",
	}

	ctx := context.Background()

	userPGRepository.EXPECT().FindByLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)

	user, err := userUC.FindByLogin(ctx, mockUser.Login)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.Login, mockUser.Login)
}

func TestUserUseCase_FindById(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userPGRepository := mock.NewMockUserPGRepository(ctrl)
	userRedisRepository := mock.NewMockUserRedisRepository(ctrl)
	apiLogger := logger.NewAPILogger(nil)
	userUC := NewUserUseCase(apiLogger, userPGRepository, userRedisRepository)

	userID := uuid.New()
	mockUser := &entity.User{
		UserID:   userID,
		Login:    "email@gmail.com",
		Password: "123456",
	}

	ctx := context.Background()

	userRedisRepository.EXPECT().GetByIDCtx(gomock.Any(), mockUser.UserID.String()).Return(nil, redis.Nil)
	userPGRepository.EXPECT().FindByID(gomock.Any(), mockUser.UserID).Return(mockUser, nil)
	userRedisRepository.EXPECT().SetUserCtx(gomock.Any(), mockUser.UserID.String(), 3600, mockUser).Return(nil)

	user, err := userUC.FindByID(ctx, mockUser.UserID)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.UserID, mockUser.UserID)
}
