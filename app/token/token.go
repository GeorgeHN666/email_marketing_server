package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Key = "ZKAIAZKAIA256578365802435_"

func GenerateJWT(user string, code string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["issuer"] = "zkaia.com"
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires in 24 hours

	// Generate the signed token string
	tokenString, err := token.SignedString([]byte(Key + code))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ValidateToken(tokenString string, code string) (bool, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the same secret key used during token creation for validation
		return []byte(Key + code), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("token expired")
	}

	return true, nil
}
