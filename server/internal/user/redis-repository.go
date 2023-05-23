package user

import (
	"context"

	"github.com/22Fariz22/passbook/server/internal/entity"
)

// UserRedisRepository Auth Redis repository interface
type UserRedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*entity.User, error)
	SetUserCtx(ctx context.Context, key string, seconds int, user *entity.User) error
	DeleteUserCtx(ctx context.Context, key string) error
}
