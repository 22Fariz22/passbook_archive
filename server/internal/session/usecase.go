package session

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
)

//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock

// SessionUseCase Session UseCase
type SessionUseCase interface {
	CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
