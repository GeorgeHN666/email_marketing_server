package serializers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	cost := 8

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func DeHash(hash string, value string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
