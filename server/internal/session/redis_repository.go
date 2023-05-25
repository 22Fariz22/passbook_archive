package session

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
)

//go:generate mockgen -source redis_repository.go -destination mock/redis_repository.go -package mock

// SessRepository Session repository
type SessRepository interface {
	CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
