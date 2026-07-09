package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	SecretKey string
}

var (
	ErrWrongTokenType = errors.New("wrong token type")
	ErrInvalidClaims  = errors.New("invalid token claims")
)

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{SecretKey: secretKey}
}

func (j *JWTManager) GenerateToken(userID string) (accessToken string, refreshToken string, err error) {
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
		"type":    "access",
	})

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"type":    "refresh",
	})

	accessToken, err = accessTokenObj.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err = refreshTokenObj.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (j *JWTManager) ValidateToken(tokenString string, expectedType string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return "", err // imza bozuk, süre geçmiş vs. — Parse zaten ayrıntılı hata döner
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidClaims // early return: sessizce aşağı düşme yok
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != expectedType {
		return "", fmt.Errorf("%w: expected %s", ErrWrongTokenType, expectedType)
	}

	userID, ok := claims["user_id"].(string)
	if !ok || userID == "" {
		return "", ErrInvalidClaims // unchecked assertion yok — panic riski kapandı
	}

	return userID, nil
}
