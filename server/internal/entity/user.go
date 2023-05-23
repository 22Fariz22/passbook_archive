package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User base model
type User struct {
	UserID   uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Login    string    `json:"login" db:"login" validate:"required,lte=30"`
	Password string    `json:"password" db:"password"`
}

// SanitizePassword Sanitize password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// HashPassword Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePasswords Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// PrepareCreate Prepare user for register
func (u *User) PrepareCreate() error {
	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

// Account entity of account data
type Account struct {
	UserID   string `json:"user_id" db:"user_id"`
	Title    string `json:"title" db:"title" validate:"required,lte=30"`
	Login    []byte `json:"login" db:"login" validate:"required,lte=30"`
	Password []byte `json:"password" db:"password" validate:"required,lte=250"`
	//CreatedAt time.Time `json:"created_at" db:"created_at"`
	//UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Text entity of text data
type Text struct {
	UserID string `json:"user_id" db:"user_id" `
	Title  string `json:"title" db:"title" validate:"required,lte=30"`
	Data   []byte `json:"data" db:"data" validate:"omitempty"`
}

// Binary entity of binary data
type Binary struct {
	UserID string `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title" validate:"required,lte=30"`
	Data   []byte `json:"data" db:"data" validate:"omitempty"`
}

// Card entity of card data
type Card struct {
	UserID     string `json:"user_id" db:"user_id"`
	Title      string `json:"title" db:"title" validate:"required,lte=30"`
	Name       []byte `db:"name"`
	CardNumber []byte `db:"card_number"`
	DateExp    []byte `db:"date_exp"`
	CVCCode    []byte `db:"cvc_code"`
}
