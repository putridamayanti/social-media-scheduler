package queue

import (
	"context"
	"github.com/redis/go-redis/v9"
	"social-media-scheduler/internal/models"
	"strconv"
)

type Scheduler struct {
	redisDb *redis.Client
}

func NewRedisClient(address string, password string, db string) *redis.Client {
	redisDb, _ := strconv.Atoi(db)
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       redisDb,
	})
}

func NewScheduler(rdb *redis.Client) *Scheduler {
	return &Scheduler{
		redisDb: rdb,
	}
}

func (s *Scheduler) AddPostQueue(ctx context.Context, post *models.Post) error {
	return s.redisDb.ZAdd(ctx, "scheduled_posts", redis.Z{
		Score:  float64(post.ScheduledAt.Unix()),
		Member: post.ID.String(),
	}).Err()
}

func (s *Scheduler) RemovePostQueue(ctx context.Context, postId string) error {
	return s.redisDb.ZRem(ctx, "scheduled_posts", postId).Err()
}
