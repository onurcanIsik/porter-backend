package service

import (
	"porter/models"
	apprand "porter/pkg/crypto"
	"porter/pkg/jwt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	userRepo         models.UserModelRepository
	tokenManager     *jwt.JWTManager
	refreshTokenRepo models.RefreshTokenModelRepository
}

func NewUserService(userRepo models.UserModelRepository, tokenManager *jwt.JWTManager, refreshTokenRepo models.RefreshTokenModelRepository) *UserService {
	return &UserService{
		userRepo:         userRepo,
		tokenManager:     tokenManager,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *UserService) LoginOrRegister(req *models.UserModel) (accessToken string, refreshToken string, err error) {
	existingUser, err := s.userRepo.GetUserByMail(req.UserMail)
	if err != nil {
		return "", "", err
	}
	if existingUser != nil {

		accessToken, refreshToken, err = s.issueTokens(existingUser.ID.String())
		if err != nil {
			return "", "", err
		}
		return accessToken, refreshToken, nil
	}

	newUser := &models.UserModel{
		ID:             uuid.New(),
		UserMail:       req.UserMail,
		UserName:       req.UserName,
		UserProfileUrl: req.UserProfileUrl,
		UserJobTitle:   req.UserJobTitle,
		UserDeviceId:   req.UserDeviceId,
		UserCreatedAt:  req.UserCreatedAt,
		UserUpdatedAt:  req.UserUpdatedAt,
		Provider:       req.Provider,
		ProviderId:     req.ProviderId,
	}

	if err = s.userRepo.CreateUser(newUser); err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err = s.issueTokens(newUser.ID.String())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) issueTokens(userID string) (string, string, error) {
	accessToken, refreshToken, err := s.tokenManager.GenerateToken(userID)
	if err != nil {
		return "", "", err
	}

	hash := apprand.HashToken(refreshToken)

	refreshTokenModel := &models.RefreshTokenModel{
		ID:           uuid.New(),
		UserID:       uuid.MustParse(userID),
		RefreshToken: hash,
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	if err := s.refreshTokenRepo.SetRefreshToken(refreshTokenModel); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) RefreshTokens(refreshToken string) (string, string, error) {
	userID, err := s.tokenManager.ValidateToken(refreshToken, "refresh")
	if err != nil {
		return "", "", err
	}

	hash := apprand.HashToken(refreshToken)

	storedToken, err := s.refreshTokenRepo.GetByTokenHash(hash)
	if err != nil {
		return "", "", err
	}

	if storedToken == nil {
		return "", "", pgx.ErrNoRows
	}

	err = s.refreshTokenRepo.DeleteByTokenHash(hash)
	if err != nil {
		return "", "", err
	}

	newTokens, newRefreshToken, err := s.issueTokens(userID)
	if err != nil {
		return "", "", err
	}

	return newTokens, newRefreshToken, nil
}
