package apprand

import (
	"crypto/rand"
	"crypto/sha256"
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

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(sum[:])
}
