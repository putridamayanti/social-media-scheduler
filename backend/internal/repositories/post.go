package repositories

import (
	"context"
	"gorm.io/gorm"
	"social-media-scheduler/internal/models"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, p *models.Post) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r *PostRepository) GetAll(ctx context.Context, query models.PostQuery) ([]models.Post, error) {
	q := r.db.WithContext(ctx)

	if query.UserId != "" {
		q = q.Where("user_id = ?", query.UserId)
	}

	var posts []models.Post
	err := q.Find(&posts).Error
	return posts, err
}

func (r *PostRepository) GetById(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error
	return &post, err
}

func (r *PostRepository) Update(ctx context.Context, id string, u map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&models.Post{}).Where("id = ?", id).Updates(u).Error
}

func (r *PostRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Post{}).Error
}
