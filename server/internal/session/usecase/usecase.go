package usecase

import (
	"context"
	"github.com/22Fariz22/passbook/server/config"
	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/22Fariz22/passbook/server/internal/session"
)

// Session use case
type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

// New session use case constructor
func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.SessionUseCase {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}

// Create new session
func (u *sessionUC) CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error) {
	return u.sessionRepo.CreateSession(ctx, session, expire)
}

// Delete session by id
func (u *sessionUC) DeleteByID(ctx context.Context, sessionID string) error {
	return u.sessionRepo.DeleteByID(ctx, sessionID)
}

// get session by id
func (u *sessionUC) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error) {
	return u.sessionRepo.GetSessionByID(ctx, sessionID)
}
