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

func (r *UserRepository) AddAccount(ctx context.Context, userID string, request *userService.AddAccountRequest) error {
	err, _ := r.db.ExecContext(ctx, addAccountQuery, userID, request.GetTitle(), request.GetLogin(), request.GetPassword())
	if err != nil {
		log.Println("err repo AddAccount in r.db.ExecContext", err)
	}
	return nil
}

func (r *UserRepository) AddText(ctx context.Context, userID string, request *userService.AddTextRequest) error {
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
	var everythingByTitle []string

	accounts := []entity.Account{}
	texts := []entity.Text{}
	cards := []entity.Card{}
	binaries := []entity.Binary{}

	//get accounts
	err := r.db.Select(&accounts, getByTitleAccountsQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByTitle:", err)
	}
	for _, v := range accounts {
		everythingByTitle = append(everythingByTitle, "account-> "+"login: "+v.Login+" "+" password: "+v.Password)
	}

	//get texts
	err = r.db.Select(&texts, getByTitleTextQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByText:", err)
	}
	for _, v := range texts {
		everythingByTitle = append(everythingByTitle, "Text-> "+v.Data)
	}

	//get cards
	err = r.db.Select(&cards, getByTitleCardQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByCard:", err)
	}
	for _, v := range cards {
		everythingByTitle = append(everythingByTitle, "Card-> "+"card number:"+v.CardNumber+" name:"+
			v.Name+" date expiration:"+v.DateExp+" cvc code:"+v.CVCCode)
	}

	//get binaries
	err = r.db.Select(&binaries, getByTitleBinaryQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByBinary:", err)
	}
	for _, v := range binaries {
		everythingByTitle = append(everythingByTitle, "Binary-> "+string(v.Data))
	}

	return everythingByTitle, nil
}

func (r *UserRepository) GetFullList(ctx context.Context, userID string) ([]string, error) {
	var everythingFullList []string

	accounts := []entity.Account{}
	texts := []entity.Text{}
	cards := []entity.Card{}
	binaries := []entity.Binary{}

	//get accounts
	err := r.db.Select(&accounts, getByFullListAccountsQuery, userID)
	if err != nil {
		log.Println("err getByFullListAccountsQuery:", err)
	}
	for _, v := range accounts {
		everythingFullList = append(everythingFullList, "account-> "+"title:"+v.Title+" login:"+v.Login+" "+" password:"+v.Password)
	}

	//get texts
	err = r.db.Select(&texts, getByFullListTextQuery, userID)
	if err != nil {
		log.Println("err getByFullListTextQuery:", err)
	}
	for _, v := range texts {
		everythingFullList = append(everythingFullList, "Text-> "+"title:"+v.Title+" data:"+v.Data)
	}

	//get cards
	err = r.db.Select(&cards, getByFullListCardQuery, userID)
	if err != nil {
		log.Println("err getByFullListCardQuery:", err)
	}
	for _, v := range cards {
		everythingFullList = append(everythingFullList, "Card-> "+"title:"+v.Title+" card number:"+v.CardNumber+" name:"+
			v.Name+" date expiration:"+v.DateExp+" cvc code:"+v.CVCCode)
	}

	//get binaries
	err = r.db.Select(&binaries, getByFullListBinaryQuery, userID)
	if err != nil {
		log.Println("err getByFullListBinaryQuery:", err)
	}
	for _, v := range binaries {
		everythingFullList = append(everythingFullList, "Binary-> "+"title:"+v.Title+"data:"+string(v.Data))
	}

	return everythingFullList, nil
}
