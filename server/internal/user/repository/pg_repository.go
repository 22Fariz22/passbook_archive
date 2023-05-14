package repository

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// User repository
type UserRepository struct {
	db *sqlx.DB
}

// User repository constructor
func NewUserPGRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create new user
func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	createdUser := &entity.User{}
	if err := r.db.QueryRowxContext(
		ctx,
		createUserQuery,
		user.Login,
		user.Password,
	).StructScan(createdUser); err != nil {
		return nil, errors.Wrap(err, "Create.QueryRowxContext")
	}

	return createdUser, nil
}

// Find by user email address
func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.GetContext(ctx, user, findByLoginQuery, login); err != nil {
		return nil, errors.Wrap(err, "FindByLogin.GetContext")
	}

	return user, nil
}

// Find user by uuid
func (r *UserRepository) FindById(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.GetContext(ctx, user, findByIDQuery, userID); err != nil {
		return nil, errors.Wrap(err, "FindById.GetContext")
	}

	return user, nil
}

func (r *UserRepository) AddAccount(ctx context.Context, userID uuid.UUID, title string, data string) error {
	if _, err := r.db.NamedExecContext(ctx,
		addAccountQuery,
		map[string]interface{}{
			"user_id": userID,
			"title":   title,
			"data":    data,
		}); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddText(ctx context.Context, userID uuid.UUID, title string, data string) error {
	if _, err := r.db.NamedExecContext(ctx,
		addTextQuery,
		map[string]interface{}{
			"user_id": userID,
			"title":   title,
			"data":    data,
		}); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddBinary(ctx context.Context, userID uuid.UUID, title string, data []byte) error {
	if _, err := r.db.NamedExecContext(ctx,
		addBinaryQuery,
		map[string]interface{}{
			"user_id": userID,
			"title":   title,
			"data":    data,
		}); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddCard(ctx context.Context, userID uuid.UUID, title string, data string) error {
	if _, err := r.db.NamedExecContext(ctx,
		addCardQuery,
		map[string]interface{}{
			"user_id": userID,
			"title":   title,
			"data":    data,
		}); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByTitle(ctx context.Context, userID uuid.UUID, title string) ([]string, error) {
	var dataRows []string

	if err := r.db.SelectContext(ctx, &dataRows, getByTitleQuery, userID, title); err != nil {
		return nil, err
	}

	return dataRows, nil
}

func (r *UserRepository) GetFullList(ctx context.Context, userID uuid.UUID) ([]string, error) {
	var dataRows []string

	if err := r.db.SelectContext(ctx, &dataRows, getFullListQuery, userID); err != nil {
		return nil, err
	}

	return dataRows, nil
}

func (r *UserRepository) GetAllTitles(ctx context.Context, userID uuid.UUID) ([]string, error) {
	var dataRows []string

	if err := r.db.SelectContext(ctx, &dataRows, getAllTitlesQuery, userID); err != nil {
		return nil, err
	}

	return dataRows, nil
}
