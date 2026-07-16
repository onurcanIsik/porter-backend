package apprand

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
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

// Görseldeki gibi sadece küçük harfler ve rakamlar içeren karakter seti
const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenerateRandomBaseUrl(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}

	// Görseldeki 'pk_' önekini (prefix) ekleyerek döndürüyoruz
	return "pk_" + string(b), nil
}
