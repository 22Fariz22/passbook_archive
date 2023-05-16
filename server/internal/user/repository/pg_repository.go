package repository

import (
	"context"
	"fmt"
	"github.com/22Fariz22/passbook/server/internal/entity"
	userService "github.com/22Fariz22/passbook/server/proto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
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

func (r *UserRepository) AddAccount(ctx context.Context, userID string, request *userService.AddAccountRequest) error { //userID uuid.UUID, title string, data string) error {
	log.Println("repo AddAccount.")
	log.Println("repo request:", request)

	//не работает
	//acc := &entity.Account{
	//	UserID:   userID,
	//	Title:    request.Title,
	//	Login:    request.Login,
	//	Password: request.Password,
	//}
	//_, err := r.db.NamedExecContext(ctx, addAccountQuery, acc)
	//if err != nil {
	//	log.Println("err repo AddAccount in r.db.ExecContext", err)
	//}

	//работает
	err, _ := r.db.ExecContext(ctx, addAccountQuery, userID, request.GetTitle(), request.GetLogin(), request.GetPassword())
	if err != nil {
		log.Println("err repo AddAccount in r.db.ExecContext", err)
	}

	////работатет
	//type TitleReturning struct {
	//	TitleReturning string `json:"title" db:"title"`
	//}
	//
	//acc := &TitleReturning{}
	//if err := r.db.QueryRowxContext(
	//	ctx,
	//	addAccountQuery,
	//	userID,
	//	request.Title,
	//	request.Login,
	//	request.Password,
	//).StructScan(acc); err != nil {
	//	return errors.Wrap(err, "Create.QueryRowxContext")
	//}

	//if _, err := r.db.NamedExecContext(ctx,
	//	addAccountQuery,
	//	map[string]interface{}{
	//		"user_id":  userID,
	//		"title":    request.Title,
	//		"login":    request.Login,
	//		"password": request.Password,
	//	}); err != nil {
	//	log.Println("err repo AddAccount in r.db.NamedExecContext", err)
	//	return err
	//}

	return nil
}

func (r *UserRepository) AddText(ctx context.Context, userID string, request *userService.AddTextRequest) error {
	fmt.Println("text repo AddText")
	err, _ := r.db.ExecContext(ctx, addTextQuery, userID, request.GetTitle(), request.GetData())
	if err != nil {
		log.Println("err repo AddText in r.db.ExecContext", err)
	}
	return nil
}

func (r *UserRepository) AddBinary(ctx context.Context, userID string, request *userService.AddBinaryRequest) error {
	fmt.Println("text repo AddBinary")
	err, _ := r.db.ExecContext(ctx, addTextQuery, userID, request.Title, request.Data)
	if err != nil {
		log.Println("err repo AddBinary in r.db.ExecContext", err)
	}
	return nil
}

func (r *UserRepository) AddCard(ctx context.Context, userID string, request *userService.AddCardRequest) error {
	fmt.Println("text repo AddCard")
	err, _ := r.db.ExecContext(ctx, addCardQuery,
		userID,
		request.Title,
		request.Name,
		request.CardNumber,
		request.DateExp,
		request.CvcCode)
	if err != nil {
		log.Println("err repo AddText in r.db.ExecContext", err)
	}
	return nil
}

func (r *UserRepository) GetByTitle(ctx context.Context, userID string, request *userService.GetByTitleRequest) ([]string, error) {
	var everything []string

	accounts := []entity.Account{}
	texts := []entity.Text{}
	cards := []entity.Card{}

	err := r.db.Select(&accounts, getByTitleAccountsQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByTitle:", err)
	}
	for _, v := range accounts {
		everything = append(everything, "account-> "+"login: "+v.Login+" "+" password: "+v.Password)
	}

	err = r.db.Select(&texts, getByTitleTextQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByText:", err)
	}
	for _, v := range texts {
		everything = append(everything, "Text-> "+v.Data)
	}

	err = r.db.Select(&cards, getByTitleCardQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByCard:", err)
	}
	for _, v := range cards {
		everything = append(everything, "Card-> "+"card number: "+v.CardNumber+" name: "+
			v.Name+" date expiration: "+v.DateExp+" cvc code: "+v.CVCCode)
	}

	//fmt.Println("ev:", everything)
	return everything, nil
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
