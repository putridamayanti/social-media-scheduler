package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserId      uuid.UUID      `gorm:"type:uuid;index;not null" json:"-"`
	Title       string         `gorm:"type:text;not null" json:"title"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	Channel     string         `gorm:"type:varchar(100);not null" json:"channel"`
	ScheduledAt time.Time      `gorm:"type:timestamptz;not null" json:"scheduled_at"`
	PublishedAt time.Time      `gorm:"type:timestamptz;not null" json:"published_at,omitempty"`
	Status      string         `gorm:"type:varchar(25);index;not null;default:draft" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type PostQuery struct {
	UserId string `query:"user_id"`
	Status string `query:"status"`
}
