package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	SecretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{SecretKey: secretKey}
}

func (j *JWTManager) GenerateToken(userID string) (accessToken string, refreshToken string, err error) {
	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
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

func (j *JWTManager) ValidateToken(tokenString string) (userID string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID = claims["user_id"].(string)
	}
	return userID, nil
}
