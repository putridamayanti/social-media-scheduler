package repositories

import (
	"context"
	"gorm.io/gorm"
	"social-media-scheduler/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	return r.db.WithContext(ctx).Create(&u).Error
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Omit("password").Find(&users).Error
	return users, err
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Omit("password").First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(ctx context.Context, id string, u map[string]interface{}) error {
	_, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Where("id = ?", id).Updates(u).Error
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}
