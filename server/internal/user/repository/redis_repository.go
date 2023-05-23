package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/22Fariz22/passbook/server/pkg/grpcerrors"
	"github.com/22Fariz22/passbook/server/pkg/logger"
	"github.com/go-redis/redis/v8"
	"time"
)

// Auth redis repository
type userRedisRepo struct {
	redisClient *redis.Client
	basePrefix  string
	logger      logger.Logger
}

// NewUserRedisRepo Auth redis repository constructor
func NewUserRedisRepo(redisClient *redis.Client, logger logger.Logger) *userRedisRepo {
	return &userRedisRepo{redisClient: redisClient, basePrefix: "user:", logger: logger}
}

// GetByIDCtx Get user by id
func (r *userRedisRepo) GetByIDCtx(ctx context.Context, key string) (*entity.User, error) {
	userBytes, err := r.redisClient.Get(ctx, r.createKey(key)).Bytes()
	if err != nil {
		if err != redis.Nil {
			return nil, grpcerrors.ErrNotFound
		}
		return nil, err
	}
	user := &entity.User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, err
	}

	return user, nil
}

// SetUserCtx Cache user with duration in seconds
func (r *userRedisRepo) SetUserCtx(ctx context.Context, key string, seconds int, user *entity.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.redisClient.Set(ctx, r.createKey(key), userBytes, time.Second*time.Duration(seconds)).Err()
}

// DeleteUserCtx Delete user by key
func (r *userRedisRepo) DeleteUserCtx(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, r.createKey(key)).Err()
}

func (r *userRedisRepo) createKey(value string) string {
	return fmt.Sprintf("%s: %s", r.basePrefix, value)
}
