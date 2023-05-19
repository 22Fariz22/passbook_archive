package repository

import (
	"context"
	"github.com/22Fariz22/passbook/server/internal/entity"
	userService "github.com/22Fariz22/passbook/server/proto"
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

func TestUserRepository_AddAccount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "title", "login", "password"}
	userUUID := uuid.New()
	mockAccount := &userService.AddAccountRequest{
		Title:    "vk",
		Login:    "email@gmail.com",
		Password: "123456",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockAccount.Title,
		mockAccount.Login,
		mockAccount.Password,
	)

	mock.ExpectQuery(addAccountQuery).WithArgs(
		mockAccount.Title,
		mockAccount.Login,
		mockAccount.Password,
	).WillReturnRows(rows)

	err = userPGRepository.AddAccount(context.Background(), userUUID.String(), mockAccount)
	require.NoError(t, err)
}

func TestUserRepository_AddText(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "title", "data"}
	userUUID := uuid.New()
	mockText := &userService.AddTextRequest{
		Title: "some text",
		Data:  "hello world",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockText.Title,
		mockText.Data,
	)

	mock.ExpectQuery(addTextQuery).WithArgs(
		mockText.Title,
		mockText.Data,
	).WillReturnRows(rows)

	err = userPGRepository.AddText(context.Background(), userUUID.String(), mockText)
	require.NoError(t, err)
}

func TestUserRepository_AddBinary(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "title", "data"}
	userUUID := uuid.New()
	mockBinary := &userService.AddBinaryRequest{
		Title: "some binary",
		Data:  []byte("hello world"),
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockBinary.Title,
		mockBinary.Data,
	)

	mock.ExpectQuery(addBinaryQuery).WithArgs(
		mockBinary.Title,
		mockBinary.Data,
	).WillReturnRows(rows)

	err = userPGRepository.AddBinary(context.Background(), userUUID.String(), mockBinary)
	require.NoError(t, err)
}

func TestUserRepository_AddCard(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	userPGRepository := NewUserPGRepository(sqlxDB)

	columns := []string{"user_id", "title", "name", "card_number", "date_exp", "cvc_code"}
	userUUID := uuid.New()
	mockCard := &userService.AddCardRequest{
		Title:      "some binary",
		Name:       "Bobik",
		CardNumber: "2333 4533 4323 4221",
		DateExp:    "11/25",
		CvcCode:    "123",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockCard.Title,
		mockCard.Name,
		mockCard.CardNumber,
		mockCard.DateExp,
		mockCard.CvcCode,
	)

	mock.ExpectQuery(addCardQuery).WithArgs(
		mockCard.Title,
		mockCard.Title,
		mockCard.Name,
		mockCard.CardNumber,
		mockCard.DateExp,
		mockCard.CvcCode,
	).WillReturnRows(rows)

	err = userPGRepository.AddCard(context.Background(), userUUID.String(), mockCard)
	require.NoError(t, err)
}
