package service

import (
	"porter/models"
	"porter/pkg/jwt"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo     models.UserModelRepository
	tokenManager *jwt.JWTManager
}

func NewUserService(userRepo models.UserModelRepository, tokenManager *jwt.JWTManager) *UserService {
	return &UserService{
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (s *UserService) LoginOrRegister(req *models.UserModelCreate) (accessToken string, refreshToken string, err error) {
	existingUser, err := s.userRepo.GetUserByMail(req.UserMail)
	if err != nil {
		return "", "", err
	}
	if existingUser != nil {
		accessToken, refreshToken, err = s.tokenManager.GenerateToken(existingUser.ID.String())
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
		UserTokenCount: req.UserTokenCount,
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

	accessToken, refreshToken, err = s.tokenManager.GenerateToken(newUser.ID.String())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
