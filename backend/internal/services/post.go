package services

import (
	"context"
	"errors"
	"social-media-scheduler/internal/models"
	"social-media-scheduler/internal/repositories"
	"strings"
	"time"
)

type PostService struct {
	postRepository repositories.PostRepositoryInterface
}

func NewPostService(postRepo repositories.PostRepositoryInterface) *PostService {
	return &PostService{postRepository: postRepo}
}

func (s *PostService) CreatePost(ctx context.Context, post *models.Post) error {
	if strings.TrimSpace(post.Content) == "" {
		return errors.New("post content is empty")
	}

	if post.ScheduledAt.Before(time.Now()) {
		return errors.New("post scheduled at is before current time")
	}

	err := s.postRepository.Create(ctx, post)
	return err
}

func (s *PostService) GetAllPosts(ctx context.Context, query models.PostQuery) ([]models.Post, error) {
	posts, err := s.postRepository.GetAll(ctx, query)
	return posts, err
}

func (s *PostService) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	post, err := s.postRepository.GetById(ctx, id)
	return post, err
}

func (s *PostService) UpdatePost(ctx context.Context, id string, u map[string]interface{}) error {
	_, err := s.postRepository.GetById(ctx, id)
	if err != nil {
		return err
	}

	err = s.postRepository.Update(ctx, id, u)
	return err
}

func (s *PostService) DeletePost(ctx context.Context, id string) error {
	err := s.postRepository.Delete(ctx, id)
	return err
}
