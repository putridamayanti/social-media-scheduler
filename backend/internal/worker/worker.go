package worker

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"social-media-scheduler/internal/models"
	"time"
)

type Worker struct {
	redisDb *redis.Client
	db      *gorm.DB
}

func NewWorker(rdb *redis.Client, db *gorm.DB) *Worker {
	return &Worker{
		redisDb: rdb,
		db:      db,
	}
}

func (w *Worker) PublishPost(ctx context.Context, postId string) error {
	var post *models.Post
	_ = w.db.WithContext(ctx).Where("id = ?", postId).First(&post)
	if post == nil {
		return errors.New("post not found")
	}

	if post.Status != "published" {
		updates := map[string]interface{}{
			"status":       "published",
			"published_at": time.Now(),
			"updated_at":   time.Now(),
		}
		return w.db.WithContext(ctx).Model(&models.Post{}).Where("id = ?", postId).Updates(updates).Error
	}

	return nil
}

func (w *Worker) Run(ctx context.Context) error {
	postIds, err := w.redisDb.ZRangeByScore(
		ctx,
		"scheduled_posts",
		&redis.ZRangeBy{
			Min:   "-inf",
			Max:   fmt.Sprintf("%d", time.Now().Unix()),
			Count: 10,
		},
	).Result()

	if err != nil {
		return err
	}

	for _, postId := range postIds {
		log.Printf("Post ID: %s\n", postId)

		locked, err := w.AcquireLock(ctx, postId)
		if err != nil {
			log.Printf("Error acquiring lock for post ID: %s %s\n", postId, err)
			continue
		}
		if !locked {
			log.Printf("Skipping lock for: %s \n", postId)
			continue
		}

		err = w.PublishPost(ctx, postId)
		if err != nil {
			log.Println("Error publishing post: "+postId, err)
			continue
		}
		log.Println("Published post: " + postId)
		res, err := w.redisDb.ZRem(ctx, "scheduled_posts", postId).Result()
		if err != nil {
			log.Println("Error removing post: "+postId, err)
			continue
		}
		log.Println("Removed post: "+postId, res)
	}

	return nil
}
