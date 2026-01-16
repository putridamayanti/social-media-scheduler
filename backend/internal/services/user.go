package services

import (
	"context"
	"errors"
	"social-media-scheduler/internal/models"
	"social-media-scheduler/internal/repositories"
	"strings"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("User email is empty")
	}

	err := s.userRepo.Create(ctx, user)
	return err
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	Users, err := s.userRepo.GetAll(ctx)
	return Users, err
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	User, err := s.userRepo.GetById(ctx, id)
	return User, err
}

func (s *UserService) UpdateUser(ctx context.Context, id string, u map[string]interface{}) error {
	_, err := s.userRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	err = s.userRepo.Update(ctx, id, u)
	return err
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	err := s.userRepo.Delete(ctx, id)
	return err
}
