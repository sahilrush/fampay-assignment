package models

import (
	"time"
)

// response struct
type Video struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Thumbnails  string    `json:"thumbnails"`
}
