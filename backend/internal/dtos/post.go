package dtos

import (
	"time"
)

type CreatePostRequest struct {
	Title       string     `json:"title" binding:"required"`
	Content     string     `json:"content" binding:"required"`
	Channel     string     `json:"channel" binding:"required"`
	Status      string     `json:"status"`
	ScheduledAt *time.Time `json:"scheduled_at"`
}

type UpdatePostRequest struct {
	Title       *string    `json:"title"`
	Content     *string    `json:"content"`
	Status      *string    `json:"status"`
	ScheduledAt *time.Time `json:"scheduled_at"`
}
