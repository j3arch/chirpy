package auth

import (
	"time"
	"uuid"

	"github.com/alexedwards/argon2id"
)

type TokenType string

const (
	TokenTypeAccess TokenType = "chirpy-access"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {

}
