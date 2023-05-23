package session

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
)

// SessRepository Session repository
type SessRepository interface {
	CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
