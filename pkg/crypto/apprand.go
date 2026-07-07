package apprand

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandomString(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomSafeString(length int) (string, error) {
	b, err := generateRandomString(length)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}
