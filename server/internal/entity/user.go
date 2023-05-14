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

// Sanitize password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Prepare user for register
func (u *User) PrepareCreate() error {
	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

// Account
type Account struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Title  string    `json:"title" db:"title" validate:"required,lte=30"`
	Data   string    `json:"data" db:"data" validate:"omitempty"`
	//CreatedAt time.Time `json:"created_at" db:"created_at"`
	//UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Text
type Text struct {
	UserID uuid.UUID `json:"user_id" db:"user_id" `
	Title  string    `json:"title" db:"title" validate:"required,lte=30"`
	Data   string    `json:"data" db:"data" validate:"omitempty"`
}

// Binary
type Binary struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Title  string    `json:"title" db:"title" validate:"required,lte=30"`
	Data   []byte    `json:"data" db:"data" validate:"omitempty"`
}

// Card
type Card struct {
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Title  string    `json:"title" db:"title" validate:"required,lte=30"`
	Data   string    `json:"data" db:"data" validate:"omitempty"`
}
