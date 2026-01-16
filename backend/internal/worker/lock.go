package worker

import (
	"context"
	"time"
)

func (w *Worker) AcquireLock(ctx context.Context, postId string) (bool, error) {
	key := "lock:post:" + postId
	return w.redisDb.SetNX(ctx, key, "1", time.Hour).Result()
}
