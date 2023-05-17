package user

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
	userService "github.com/22Fariz22/passbook/server/proto"
	"github.com/google/uuid"
)

// User UseCase interface
type UserUseCase interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, email string, password string) (*entity.User, error)
	FindByLogin(ctx context.Context, login string) (*entity.User, error)
	FindById(ctx context.Context, userID uuid.UUID) (*entity.User, error)

	AddAccount(ctx context.Context, userID string, request *userService.AddAccountRequest) error //userID uuid.UUID, tittle string, data string) error
	AddText(ctx context.Context, userID string, request *userService.AddTextRequest) error
	AddBinary(ctx context.Context, userID string, request *userService.AddBinaryRequest) error
	AddCard(ctx context.Context, userID string, request *userService.AddCardRequest) error
	GetByTitle(ctx context.Context, userID string, request *userService.GetByTitleRequest) ([]string, error)
	GetFullList(ctx context.Context, userID string) ([]string, error)
}
