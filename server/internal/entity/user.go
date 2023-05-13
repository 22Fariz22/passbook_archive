package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User base model
type User struct {
	UserID   uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Login    string    `json:"login" db:"login" validate:"required,lte=30"`
	Password string    `json:"password,omitempty" db:"password"`
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
	ID        string
	Login     string
	Password  string
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Text
type Text struct {
	ID        string
	TextData  string
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Binary
type Binary struct {
	ID         string
	BinaryData byte
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Card
type Card struct {
	ID             string
	UploadedAt     time.Time
	CardNumber     string
	ExpirationData string
	CardHolderName string
	CVCCode        string
	CreatedAt      time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
