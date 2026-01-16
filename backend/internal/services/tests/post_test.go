package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"social-media-scheduler/internal/models"
	"social-media-scheduler/internal/services"
	"testing"
	"time"
)

type MockPostRepo struct {
	mock.Mock
}

func (repo *MockPostRepo) Create(ctx context.Context, p *models.Post) error {
	args := repo.Called(ctx, p)
	return args.Error(0)
}

func (repo *MockPostRepo) GetAll(ctx context.Context, query models.PostQuery) ([]models.Post, error) {
	args := repo.Called(ctx, query)
	return args.Get(0).([]models.Post), args.Error(1)
}

func (repo *MockPostRepo) GetById(ctx context.Context, id string) (*models.Post, error) {
	args := repo.Called(ctx, id)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (repo *MockPostRepo) Update(ctx context.Context, id string, u map[string]interface{}) error {
	args := repo.Called(ctx, id, u)
	return args.Error(0)
}

func (repo *MockPostRepo) Delete(ctx context.Context, id string) error {
	args := repo.Called(ctx, id)
	return args.Error(0)
}

func TestPostServiceCreate_Valid(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	service := services.NewPostService(mockPostRepo)

	post := models.Post{
		Content:     "test post",
		ScheduledAt: time.Now().Add(1 * time.Hour),
	}

	mockPostRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	err := service.CreatePost(context.Background(), &post)

	assert.NoError(t, err)
	assert.Equal(t, post.Content, "test post")
}

func TestPostServiceCreate_EmptyContent(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	service := services.NewPostService(mockPostRepo)

	post := models.Post{
		Content:     "",
		ScheduledAt: time.Now().Add(1 * time.Hour),
	}

	mockPostRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	err := service.CreatePost(context.Background(), &post)

	assert.Error(t, err)
	assert.Equal(t, "post content is empty", err.Error())
}

func TestPostServiceCreate_InvalidScheduledTime(t *testing.T) {
	mockPostRepo := new(MockPostRepo)
	service := services.NewPostService(mockPostRepo)

	post := models.Post{
		Content:     "test post",
		ScheduledAt: time.Now().Add(-1 * time.Hour),
	}

	mockPostRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	err := service.CreatePost(context.Background(), &post)

	assert.Error(t, err)
	assert.Equal(t, "post scheduled at is before current time", err.Error())
}
