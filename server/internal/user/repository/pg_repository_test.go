package repository

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "login", "password"}
	userUUID := uuid.New()
	mockUser := &entity.User{
		UserID:   userUUID,
		Login:    "email@gmail.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockUser.Login,
		mockUser.Password,
	)

	mock.ExpectQuery(createUserQuery).WithArgs(
		mockUser.Login,
		mockUser.Password,
	).WillReturnRows(rows)

	createdUser, err := userPGRepository.Create(context.Background(), mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
}

func TestUserRepository_FindBylogin(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "login", "password"}
	userUUID := uuid.New()
	mockUser := &entity.User{
		UserID:   userUUID,
		Password: "123456",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockUser.Login,
		mockUser.Password,
	)

	mock.ExpectQuery(findByLoginQuery).WithArgs(mockUser.Login).WillReturnRows(rows)

	foundUser, err := userPGRepository.FindByLogin(context.Background(), mockUser.Login)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.Login, mockUser.Login)
}

func TestUserRepository_FindById(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "login", "password"}
	userUUID := uuid.New()
	mockUser := &entity.User{
		UserID:   userUUID,
		Login:    "email@gmail.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockUser.Login,
		mockUser.Password,
	)

	mock.ExpectQuery(findByIDQuery).WithArgs(mockUser.UserID).WillReturnRows(rows)

	foundUser, err := userPGRepository.FindById(context.Background(), mockUser.UserID)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.UserID, mockUser.UserID)
}
