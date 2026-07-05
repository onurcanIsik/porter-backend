package service

import (
	"fmt"
	"porter/models"
	"porter/pkg/jwt"
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

func (s *UserService) RegisterUser(req *models.UserModelCreate) (*models.UserModel, error) {
	existingUser, err := s.userRepo.GetUserByMail(req.UserMail)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.UserMail)
	}

	return nil, nil
}
