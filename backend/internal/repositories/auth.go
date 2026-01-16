package repositories

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"social-media-scheduler/internal/models"
	"time"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateSession(ctx context.Context, userId uuid.UUID) (*models.Session, error) {
	session := &models.Session{
		ID:        uuid.New(),
		UserID:    userId,
		ExpiresAt: time.Now().Add(time.Hour * 24),
		CreatedAt: time.Now(),
	}

	err := r.db.WithContext(ctx).Create(session).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *AuthRepository) GetSession(ctx context.Context, id uuid.UUID) (*models.Session, error) {
	var session models.Session
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *AuthRepository) RemoveSession(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Session{}).Error
}
