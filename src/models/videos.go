package models

import "time"

type Video struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Thumbnails  string    `json:"thumbnails"`
}
