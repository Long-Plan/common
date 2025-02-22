package hash

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type hashPassword struct {
}

func NewHashPassword() IHashPassword {
	return &hashPassword{}
}

type IHashPassword interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword, password string) error
}

// HashPassword creates a hashed version of a plaintext password.
func (hp *hashPassword) HashPassword(password string) (string, error) {
	if len(password) > 72 {
		h := sha256.Sum256([]byte(password))
		password = fmt.Sprintf("%x", h)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plaintext password.
func (hp *hashPassword) ComparePasswords(hashedPassword, password string) error {
	if len(password) > 72 {
		h := sha256.Sum256([]byte(password))
		password = fmt.Sprintf("%x", h)
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
