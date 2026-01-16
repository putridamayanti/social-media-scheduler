package services

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"social-media-scheduler/internal/dtos"
	"social-media-scheduler/internal/models"
	"social-media-scheduler/internal/repositories"
)

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
	authRepo repositories.AuthRepositoryInterface
}

func NewAuthService(authRepo repositories.AuthRepositoryInterface, userRepo repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{authRepo: authRepo, userRepo: userRepo}
}

func (s *AuthService) Login(ctx context.Context, request dtos.LoginRequest) (*models.Session, error) {
	user, err := s.userRepo.GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	session, err := s.authRepo.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *AuthService) Register(ctx context.Context, request dtos.RegisterRequest) error {
	user, _ := s.userRepo.GetByEmail(ctx, request.Email)
	if user.Email != "" {
		return errors.New("User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userCreate := &models.User{
		Email:    request.Email,
		Password: string(hash),
		Name:     request.Name,
	}

	err = s.userRepo.Create(ctx, userCreate)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Logout(ctx context.Context, userId string) error {
	return s.authRepo.RemoveSession(ctx, userId)
}
