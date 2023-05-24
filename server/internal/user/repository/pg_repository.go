package repository

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"

	"github.com/22Fariz22/passbook/cli/config"
	"github.com/22Fariz22/passbook/server/internal/entity"
	userService "github.com/22Fariz22/passbook/server/proto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// UserRepository User repository
type UserRepository struct {
	db *sqlx.DB
}

// NewUserPGRepository User repository constructor
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

// FindByLogin Find by user email address
func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.GetContext(ctx, user, findByLoginQuery, login); err != nil {
		return nil, errors.Wrap(err, "FindByLogin.GetContext")
	}

	return user, nil
}

// FindByID Find user by uuid
func (r *UserRepository) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.GetContext(ctx, user, findByIDQuery, userID); err != nil {
		return nil, errors.Wrap(err, "FindByID.GetContext")
	}

	return user, nil
}

// AddAccount Add  account data
func (r *UserRepository) AddAccount(ctx context.Context, userID string, request *userService.AddAccountRequest) error {
	encLogin := encrypt(request.Login)
	encPassword := encrypt(request.Password)

	_, err := r.db.ExecContext(ctx, addAccountQuery, userID, request.GetTitle(), encLogin, encPassword)
	if err != nil {
		log.Println("err repo AddAccount in r.db.ExecContext", err)
	}
	return nil
}

// AddText add text data
func (r *UserRepository) AddText(ctx context.Context, userID string, request *userService.AddTextRequest) error {
	encData := encrypt(request.Data)

	_, err := r.db.ExecContext(ctx, addTextQuery, userID, request.GetTitle(), encData)
	if err != nil {
		log.Println("err repo AddText in r.db.ExecContext", err)
	}
	return nil
}

// AddBinary add binary data
func (r *UserRepository) AddBinary(ctx context.Context, userID string, request *userService.AddBinaryRequest) error {
	_, err := r.db.ExecContext(ctx, addBinaryQuery, userID, request.Title, request.Data)
	if err != nil {
		log.Println("err repo AddBinary in r.db.ExecContext", err)
	}
	return nil
}

// AddCard add card data
func (r *UserRepository) AddCard(ctx context.Context, userID string, request *userService.AddCardRequest) error {
	encCardNumber := encrypt(request.CardNumber)
	encName := encrypt(request.Name)
	encDateExp := encrypt(request.DateExp)
	encCVCCode := encrypt(request.CvcCode)

	err, _ := r.db.ExecContext(ctx, addCardQuery,
		userID,
		request.Title,
		encName,
		encCardNumber,
		encDateExp,
		encCVCCode)
	if err != nil {
		log.Println("err repo AddText in r.db.ExecContext", err)
	}
	return nil
}

// GetByTitle find data by title
func (r *UserRepository) GetByTitle(ctx context.Context, userID string, request *userService.GetByTitleRequest) ([]string, error) {
	//here all data from db
	var everythingByTitle []string

	accounts := []entity.Account{}
	texts := []entity.Text{}
	cards := []entity.Card{}
	binaries := []entity.Binary{}

	//get accounts
	err := r.db.SelectContext(ctx, &accounts, getByTitleAccountsQuery, userID, request.Title)
	if err != nil {
		log.Println("err in repo GetByAccount:", err)
	}
	for _, v := range accounts {
		decrLogin := decrypt(v.Login)
		decrPassword := decrypt(v.Password)
		everythingByTitle = append(everythingByTitle, "account-> "+"login: "+decrLogin+" "+" password: "+decrPassword)
	}

	//get texts
	err = r.db.SelectContext(ctx, &texts, getByTitleTextQuery, userID, request.Title)
	if err != nil {
		log.Println("err in repo err GetByText:", err)
	}

	for _, v := range texts {
		decrData := decrypt(v.Data)
		everythingByTitle = append(everythingByTitle, "Text-> "+decrData)
	}

	//get cards
	err = r.db.SelectContext(ctx, &cards, getByTitleCardQuery, userID, request.Title)
	if err != nil {
		log.Println("err in repo  GetByCard:", err)
	}
	for _, v := range cards {
		decrName := decrypt(v.Name)
		decrCardNumber := decrypt(v.CardNumber)
		decrDateExp := decrypt(v.DateExp)
		decrCVCCode := decrypt(v.CVCCode)

		everythingByTitle = append(everythingByTitle, "Card-> "+"card number:"+decrCardNumber+" name:"+
			decrName+" date expiration:"+decrDateExp+" cvc code:"+decrCVCCode)
	}

	//get binaries
	err = r.db.SelectContext(ctx, &binaries, getByTitleBinaryQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByBinary:", err)
	}
	for _, v := range binaries {
		everythingByTitle = append(everythingByTitle, "Binary-> "+string(v.Data))
	}

	return everythingByTitle, nil
}

// GetFullList find all type of data
func (r *UserRepository) GetFullList(ctx context.Context, userID string) ([]string, error) {
	//here all data
	var everythingFullList []string

	accounts := []entity.Account{}
	texts := []entity.Text{}
	cards := []entity.Card{}
	binaries := []entity.Binary{}

	//get accounts
	err := r.db.SelectContext(ctx, &accounts, getByFullListAccountsQuery, userID)
	if err != nil {
		log.Println("err getByFullListAccountsQuery:", err)
	}

	for _, v := range accounts {
		decrLogin := decrypt(v.Login)
		decrPassword := decrypt(v.Password)
		everythingFullList = append(everythingFullList, "account-> "+"title:"+v.Title+" login:"+decrLogin+" "+" password:"+decrPassword)
	}

	//get texts
	err = r.db.SelectContext(ctx, &texts, getByFullListTextQuery, userID)
	if err != nil {
		log.Println("err getByFullListTextQuery:", err)
		return nil, err
	}

	for _, v := range texts {
		decrData := decrypt(v.Data)
		everythingFullList = append(everythingFullList, "Text-> "+"title:"+v.Title+" data:"+decrData)
	}

	//get cards
	err = r.db.SelectContext(ctx, &cards, getByFullListCardQuery, userID)
	if err != nil {
		log.Println("err getByFullListCardQuery:", err)
	}

	for _, v := range cards {
		decrCardNumber := decrypt(v.CardNumber)
		decrName := decrypt(v.Name)
		decrDateExp := decrypt(v.DateExp)
		decrCVCCode := decrypt(v.CVCCode)
		everythingFullList = append(everythingFullList, "Card-> "+"title:"+v.Title+" card number:"+decrCardNumber+" name:"+
			decrName+" date expiration:"+decrDateExp+" cvc code:"+decrCVCCode)
	}

	//get binaries
	err = r.db.SelectContext(ctx, &binaries, getByFullListBinaryQuery, userID)
	if err != nil {
		log.Println("err getByFullListBinaryQuery:", err)
	}

	for _, v := range binaries {
		everythingFullList = append(everythingFullList, "Binary-> "+"title:"+v.Title+"data:"+string(v.Data))
	}

	return everythingFullList, nil
}

// encrypt data
func encrypt(s string) []byte {
	text := []byte(s)
	key := []byte(config.Key)

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		log.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err)
	}

	return gcm.Seal(nonce, nonce, text, nil)
}

// decrypt data
func decrypt(ciphertext []byte) string {
	key := []byte(config.Key)

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}

	return string(plaintext)
}
