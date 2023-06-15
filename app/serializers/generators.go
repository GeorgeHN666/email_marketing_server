package serializers

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

// GenerateCode Will generate a code to a given user
func GenerateAuthCode() (string, error) {
	length := 16

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return fmt.Sprintf("ZCS-%s", string(bytes)), nil
}

func GenerateNumCode() (int, error) {
	length := 6

	chars := "0123456789"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return 0, err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	code, _ := strconv.Atoi(string(bytes))

	return code, nil
}
