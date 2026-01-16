package repositories

import (
	"context"
	"github.com/google/uuid"
	"social-media-scheduler/internal/models"
)

type UserRepositoryInterface interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, u *models.User) error
}

type AuthRepositoryInterface interface {
	CreateSession(ctx context.Context, userId uuid.UUID) (*models.Session, error)
	RemoveSession(ctx context.Context, id string) error
}

type PostRepositoryInterface interface {
	Create(ctx context.Context, p *models.Post) error
	GetAll(ctx context.Context, query models.PostQuery) ([]models.Post, error)
	GetById(ctx context.Context, id string) (*models.Post, error)
	Update(ctx context.Context, id string, u map[string]interface{}) error
	Delete(ctx context.Context, id string) error
}
