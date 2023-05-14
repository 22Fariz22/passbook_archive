package user

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/google/uuid"
)

// User pg repository
type UserPGRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByLogin(ctx context.Context, login string) (*entity.User, error)
	FindById(ctx context.Context, userID uuid.UUID) (*entity.User, error)

	AddAccount(ctx context.Context, userID uuid.UUID, tittle string, data string) error
	AddText(ctx context.Context, userID uuid.UUID, tittle string, data string) error
	AddBinary(ctx context.Context, userID uuid.UUID, tittle string, data []byte) error
	AddCard(ctx context.Context, userID uuid.UUID, tittle string, data string) error
	GetByTitle(ctx context.Context, userID uuid.UUID, title string) ([]string, error)
	GetFullList(ctx context.Context, userID uuid.UUID) ([]string, error)
	GetAllTitles(ctx context.Context, userID uuid.UUID) ([]string, error)
}
